package healthCheck

type Response struct {
	Version string `json:"version"`
	Status  string `json:"status"`
}
