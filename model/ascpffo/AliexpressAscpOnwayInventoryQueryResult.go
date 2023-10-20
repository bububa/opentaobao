package ascpffo

// AliexpressascponwayinventoryqueryResult 结构体
type AliexpressascponwayinventoryqueryResult struct {
	// 出参列表
	DataList []ErpOnWayInventoryDto `json:"data_list,omitempty" xml:"data_list>erp_on_way_inventory_dto,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
