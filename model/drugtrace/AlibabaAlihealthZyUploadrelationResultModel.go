package drugtrace

// AlibabaAlihealthZyUploadrelationResultModel 结构体
type AlibabaAlihealthZyUploadrelationResultModel struct {
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// model
	Model *SaveCodeRelationResultDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}
