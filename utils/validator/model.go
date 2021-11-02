package validator

type ErrResponse struct {
	Errors []string `json:"errors"`
}
