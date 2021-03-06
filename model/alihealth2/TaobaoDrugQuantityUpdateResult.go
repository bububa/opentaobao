package alihealth2

// TaobaoDrugQuantityUpdateResult 结构体
type TaobaoDrugQuantityUpdateResult struct {
	// 错误编号
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 请求的接口是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
