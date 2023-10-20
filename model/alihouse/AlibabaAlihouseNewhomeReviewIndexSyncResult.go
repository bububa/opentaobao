package alihouse

// AlibabaalihousenewhomereviewindexsyncResult 结构体
type AlibabaalihousenewhomereviewindexsyncResult struct {
	// 返回信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回素材id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// true或false
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
