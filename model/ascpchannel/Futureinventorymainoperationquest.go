package ascpchannel

// Futureinventorymainoperationquest 结构体
type Futureinventorymainoperationquest struct {
	// 操作主单
	MainOrderDto *Mainorderdto `json:"main_order_dto,omitempty" xml:"main_order_dto,omitempty"`
	// 子单操作明细列表
	DetailOperationDtos []Detailoperationdtos `json:"detail_operation_dtos,omitempty" xml:"detail_operation_dtos>detailoperationdtos,omitempty"`
}
