package qimen

// TaobaoqimeninventorybatchqueryRequest 结构体
type TaobaoqimeninventorybatchqueryRequest struct {
	// 货主编码，string(50)，必填
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 仓库编码，string(50)，必填
	WarehouseCode string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
	// ERP 系统商品编码, string(50)，条件必填
	ItemCode string `json:"itemCode,omitempty" xml:"itemCode,omitempty"`
	// 仓储系统商品编码, string(50)，条件必填
	ItemId string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// 当前页，从 1 开始，必填
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页条数(最多 100 条)，必填
	PageSize int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}
