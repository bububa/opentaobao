package msgamp

// TaobaobcchatmessagesendResult 结构体
type TaobaobcchatmessagesendResult struct {
	// SERVICE_ERROR
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// model
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// SERVICE_ERROR
	MsgErrMessage string `json:"msg_err_message,omitempty" xml:"msg_err_message,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
