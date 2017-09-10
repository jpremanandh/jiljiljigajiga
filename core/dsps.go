package core

type Dsps struct {
	URL     string `json:"url"`
	ID      string `json:"id"`
	QPS     string `json:"qps"`
	limiter Limiter
}
