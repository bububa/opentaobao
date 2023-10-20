package wenyuvideo

// YoukuwenyuvideoseetagetResult 结构体
type YoukuwenyuvideoseetagetResult struct {
	// 返回数据
	Values []YoukuwenyuvideoseetagetModel `json:"values,omitempty" xml:"values>youkuwenyuvideoseetaget_model,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 结果代码
	HttpStatusCode int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
