package wdk

// TaobaoWdkEquipmentConveyorConveyorinfoGetResult 结构体
type TaobaoWdkEquipmentConveyorConveyorinfoGetResult struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// model
	Model *WcsConveyorInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
