package core

var RequestBody struct {
	Expression string `json:"expression"`
}

type ResultBody struct {
	Result float64 `json:"result"`
}

type ErrorBody struct {
	Error string `json:"error"`
}
