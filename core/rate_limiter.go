package core

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/andres-erbsen/clock"
)

type Limiter interface {
	Take() (time.Time, bool)
}

// Clock is the minimum necessary interface to instantiate a rate limiter with
// a clock or mock clock, compatible with clocks created using
// github.com/andres-erbsen/clock.
type Clock interface {
	Now() time.Time
	Sleep(time.Duration)
}

type limiter struct {
	sync.Mutex
	last       time.Time
	sleepFor   time.Duration
	perRequest time.Duration
	maxSlack   time.Duration
	clock      Clock
}

// Option configures a Limiter.
type Option func(l *limiter)

// New returns a Limiter that will limit to the given Rate.
func NewRateLimiter(rateLimit string, opts ...Option) (Limiter, error) {

	if rateLimit == "" {
		l := NewUnlimited()
		return l, nil
	}

	rateReg, _ := regexp.Compile("^[0-9]*")
	rate, convErr := strconv.ParseInt(rateReg.FindString(rateLimit), 10, 32)
	if convErr != nil {
		return nil, errors.New("wrong value for the config")
	}

	limiterReg, _ := regexp.Compile("[a-zA-Z]*$")
	limiterType := strings.ToLower(limiterReg.FindString(rateLimit))
	var request, slack time.Duration
	slackLimit := time.Duration(-1 * rate / 2)

	if rate == 0 {
		return nil, errors.New("Wrong Config")
	}

	switch limiterType {
	case "rpm":
		{
			request = time.Minute / time.Duration(rate)
			slack = slackLimit * time.Minute / time.Duration(rate)
		}
	case "rph":
		{
			request = time.Hour / time.Duration(rate)
			slack = slackLimit * time.Hour / time.Duration(rate)
		}
	case "rps":
		{
			request = time.Second / time.Duration(rate)
			slack = slackLimit * time.Second / time.Duration(rate)
		}
	default:
		{
			return nil, errors.New("Wrong Config")
		}
	}

	l := &limiter{
		perRequest: request,
		maxSlack:   slack,
	}
	for _, opt := range opts {
		opt(l)
	}
	if l.clock == nil {
		l.clock = clock.New()
	}
	return l, nil
}

// WithClock returns an option for ratelimit.New that provides an alternate
// Clock implementation, typically a mock Clock for testing.
func WithClock(clock Clock) Option {
	return func(l *limiter) {
		l.clock = clock
	}
}

// WithoutSlack is an option for ratelimit.New that initializes the limiter
// without any initial tolerance for bursts of traffic.
var WithoutSlack Option = withoutSlackOption

func withoutSlackOption(l *limiter) {
	l.maxSlack = 0
}

// Take blocks to ensure that the time spent between multiple
// Take calls is on average time.Second/rate.
func (t *limiter) Take() (time.Time, bool) {
	t.Lock()
	defer t.Unlock()

	now := t.clock.Now()

	if t.last.IsZero() {
		t.last = now
		return t.last, true
	}

	t.sleepFor += t.perRequest - now.Sub(t.last)

	if t.sleepFor < t.maxSlack {
		t.sleepFor = t.maxSlack
	}

	if t.sleepFor > 0 {
		t.clock.Sleep(t.sleepFor)
		return t.last, false
		// t.last = now.Add(t.sleepFor)
		// t.sleepFor = 0
	} else {
		t.last = now
		return t.last, true
	}
}

type unlimited struct{}

func NewUnlimited() Limiter {
	return unlimited{}
}

func (unlimited) Take() (time.Time, bool) {
	return time.Now(), true
}
