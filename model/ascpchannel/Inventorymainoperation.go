package ascpchannel

// Inventorymainoperation 结构体
type Inventorymainoperation struct {
	// 子单操作明细
	DetailOperationList []Detailoperationlist `json:"detail_operation_list,omitempty" xml:"detail_operation_list>detailoperationlist,omitempty"`
	// 操作主单
	MainOrder *Mainorder `json:"main_order,omitempty" xml:"main_order,omitempty"`
}
