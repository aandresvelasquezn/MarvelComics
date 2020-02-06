package comic

// DataWrapper Estrctura
type DataWrapper struct {
	Code  int  `json:"code"`
	Datas Data `json:"data"`
}

//Data Structure
type Data struct {
	Results []Result `json:"results"`
}

//Result structure
type Result struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
