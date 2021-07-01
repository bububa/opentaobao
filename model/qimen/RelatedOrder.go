package qimen

// RelatedOrder 结构体
type RelatedOrder struct {
	// 关联的订单类型(CG=采购单;DB=调拨单;CK=出库单;RK=入库单;只传英文编码)
	OrderType string `json:"orderType,omitempty" xml:"orderType,omitempty"`
	// 关联的订单编号
	OrderCode string `json:"orderCode,omitempty" xml:"orderCode,omitempty"`
}
