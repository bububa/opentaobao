package fenxiao

// TaobaoInventoryWarehouseManageResult 结构体
type TaobaoInventoryWarehouseManageResult struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// data
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
