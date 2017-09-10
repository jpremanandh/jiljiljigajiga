### go run jiljiljigajiga.go

Expected initial changes to be incorporated for this project

### SSP Tracker Module (CONTRIBUTOR: PREM)
1. Receive bid request (Status: COMPLETED)
2. Make async API calls to all DSPS
  * Timeout from getting the initial request to collecting annd request from DSPS and conducting auctions should be less than 150ms
  * Implement QPS control
3. Collect response and conduct auction
4. Respond to publisher with tracker

### SSP Tracker Module (CONTRIBUTOR: RAMKUMAR)
1. Expose API endpoint for ssp tracker
2. Log the query params for requset

### Request handling architecture
1. Get Request
2. Lopp through all DSPS
    1. Check if Request is eligible to process the request
    2. (Incoming request's context matches with DSPS priority) AND (QPS Frequency is within limit with the new request)