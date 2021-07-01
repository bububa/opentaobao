package tax

// ResultItem 结构体
type ResultItem struct {
	// 每一项成功失败
	Correctness bool `json:"correctness,omitempty" xml:"correctness,omitempty"`
	// 每一项具体异常信息
	ErrorDescription string `json:"error_description,omitempty" xml:"error_description,omitempty"`
	// 每一项的请求id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}
