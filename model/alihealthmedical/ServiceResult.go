package alihealthmedical

// ServiceResult 结构体
type ServiceResult struct {
	// errCode
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// errMessage
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// 审核记录通讯id
	DoctorUuid string `json:"doctor_uuid,omitempty" xml:"doctor_uuid,omitempty"`
	// 错误码
	CodeError string `json:"code_error,omitempty" xml:"code_error,omitempty"`
	// 错误信息
	MessageError string `json:"message_error,omitempty" xml:"message_error,omitempty"`
	// 返回数据对象
	Data *OuterMsgPullVo `json:"data,omitempty" xml:"data,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 消息是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
