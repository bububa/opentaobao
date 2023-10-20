package qimen

// DeliveryOrderQueryRequest 结构体
type DeliveryOrderQueryRequest struct {
	// 货主编码
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 仓库编码
	WarehouseCode string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
	// 发库单号
	OrderCode string `json:"orderCode,omitempty" xml:"orderCode,omitempty"`
	// 仓储系统发库单号
	OrderId string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// 交易单号
	OrderSourceCode string `json:"orderSourceCode,omitempty" xml:"orderSourceCode,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 当前页
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页orderLine条数(最多100条)
	PageSize int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoqimendeliveryorderqueryMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}
