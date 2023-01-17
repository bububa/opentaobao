package charity

// CsrResult 结构体
type CsrResult struct {
	// 接口响应消息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 状态码：200表示正常，非200表示异常
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 公益时授权链接结果
	Data *JumpAddressHsfResponse `json:"data,omitempty" xml:"data,omitempty"`
}
