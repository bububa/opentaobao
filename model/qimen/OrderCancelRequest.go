package qimen

// OrderCancelRequest 结构体
type OrderCancelRequest struct {
	// 交易平台子订单信息
	SubSourceOrders []SubSourceOrder `json:"subSourceOrders,omitempty" xml:"subSourceOrders>sub_source_order,omitempty"`
	// 仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
	WarehouseCode string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
	// 货主编码
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 单据编码
	OrderCode string `json:"orderCode,omitempty" xml:"orderCode,omitempty"`
	// 仓储系统单据编码
	OrderId string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// 单据类型(JYCK=一般交易出库单;HHCK= 换货出库;BFCK=补发出库;PTCK=普通出库单;DBCK=调拨出库;B2BRK=B2B入库;B2BCK=B2B出库;QTCK=其他出库;SCRK=生产入库;LYRK=领用入库;CCRK=残次品入库;CGRK=采购入库;DBRK= 调拨入库;QTRK=其他入库;XTRK= 销退入库;THRK=退货入库;HHRK= 换货入库;CNJG= 仓内加工单;CGTH=采购退货出库单)
	OrderType string `json:"orderType,omitempty" xml:"orderType,omitempty"`
	// 取消原因
	CancelReason string `json:"cancelReason,omitempty" xml:"cancelReason,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoqimenordercancelMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}
