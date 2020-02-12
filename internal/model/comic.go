package model

// DataWrapper Structure
type DataWrapper struct {
	Code int `json:"code"`
	Data struct {
		Result []struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
		} `json:"results"`
	} `json:"data"`
}
