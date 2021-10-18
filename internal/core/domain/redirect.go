package domain

type Redirect struct {
	Code string `json:"code"`
	URL  string `json:"url"`
}
