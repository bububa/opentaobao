package util

// ProcessResponse 结构体
type ProcessResponse struct {
	// 流程实例ID
	ProcessInstanceId string `json:"process_instance_id,omitempty" xml:"process_instance_id,omitempty"`
	// 流程错误码
	ProcessErrorCode string `json:"process_error_code,omitempty" xml:"process_error_code,omitempty"`
	// 流程错误信息
	ProcessErrorMsg string `json:"process_error_msg,omitempty" xml:"process_error_msg,omitempty"`
	// 流程备注
	ProcessRemark string `json:"process_remark,omitempty" xml:"process_remark,omitempty"`
	// 流程返回值
	ProcessData string `json:"process_data,omitempty" xml:"process_data,omitempty"`
	// 流程执行是否成功
	ProcessSuccess bool `json:"process_success,omitempty" xml:"process_success,omitempty"`
	// 业务是否成功
	BizSuccess bool `json:"biz_success,omitempty" xml:"biz_success,omitempty"`
}
