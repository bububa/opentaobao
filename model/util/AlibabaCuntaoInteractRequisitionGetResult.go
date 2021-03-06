package util

// AlibabaCuntaoInteractRequisitionGetResult 结构体
type AlibabaCuntaoInteractRequisitionGetResult struct {
	// 异常时返回的code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 异常时返回的描述
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// model
	Model *ListDto `json:"model,omitempty" xml:"model,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
