package wdk

// TaobaoWdkEquipmentConveyorSystemeventGetResult 结构体
type TaobaoWdkEquipmentConveyorSystemeventGetResult struct {
	// 返回的数据
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 返回值与返回的信息
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
}
