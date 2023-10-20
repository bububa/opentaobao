package alihouse

// AlibabaalihousemessageworkorderpushResult 结构体
type AlibabaalihousemessageworkorderpushResult struct {
	// 错误信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 消息记录id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
