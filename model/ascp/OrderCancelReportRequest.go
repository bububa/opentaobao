package ascp

// OrderCancelReportRequest 结构体
type OrderCancelReportRequest struct {
	// 交易平台子订单信息
	SubSourceOrders []SubSourceOrder `json:"sub_source_orders,omitempty" xml:"sub_source_orders>sub_source_order,omitempty"`
	// 仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 货主编码
	OwnerCode string `json:"owner_code,omitempty" xml:"owner_code,omitempty"`
	// 单据编码
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 仓储系统单据编码
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 单据类型(JYCK=一般交易出库单;HHCK= 换货出库;BFCK=补发出库;PTCK=普通出库单;DBCK=调拨出库;B2BRK=B2B入库;B2BCK=B2B出库;QTCK=其他出库;SCRK=生产入库;LYRK=领用入库;CCRK=残次品入库;CGRK=采购入库;DBRK= 调拨入库;QTRK=其他入库;XTRK= 销退入库;THRK=退货入库;HHRK= 换货入库;CNJG= 仓内加工单;CGTH=采购退货出库单)
	OrderType string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 取消原因
	CancelReason string `json:"cancel_reason,omitempty" xml:"cancel_reason,omitempty"`
	// 扩展属性
	ExtendProps string `json:"extend_props,omitempty" xml:"extend_props,omitempty"`
	// 业务请求ID，用于做幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 业务请求时间（时间戳）
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}
