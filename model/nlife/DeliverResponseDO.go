package nlife

// DeliverResponseDo 结构体
type DeliverResponseDo struct {
	// 零售商的门店id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 关联的门店采购订单号
	TradeNo string `json:"trade_no,omitempty" xml:"trade_no,omitempty"`
	// 供应商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 查询的总数
	TotalResults string `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 发货单列表
	DeliverList []Deliverlist `json:"deliver_list,omitempty" xml:"deliver_list>deliverlist,omitempty"`
}
