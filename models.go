package main
type IPStat struct {
		IP    string
		Count int
	}
	type EndpointStat struct {
    Endpoint string
    Count    int
}
type StatusStat struct {
    Status string
    Count  int
}

type Report struct {
	TotalRequests int
	Error404      int
	Error5xx      int
	TopIPs        []IPStat
	TopEndpoints  []EndpointStat
	StatusCodes   []StatusStat
}
