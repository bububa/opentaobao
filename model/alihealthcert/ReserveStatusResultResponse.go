package alihealthcert

// ReserveStatusResultResponse 结构体
type ReserveStatusResultResponse struct {
	// 业务响应code
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 正文
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
}
