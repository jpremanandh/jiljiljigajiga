# jiljiljigajiga

Expected initial changes to be incorporated for this project

SSP Module
-> Receive bid request
-> Make async API calls to all DSPS
    .) Timeout from getting the initial request to collecting annd request from DSPS and conducting auctions should be less than 150ms
    .) Implement QPS control
-> Collect response and conduct auction
-> Respond to publisher with tracker

SSP Tracker
-> Expose API endpoint for ssp tracker
-> Log the query params for requset