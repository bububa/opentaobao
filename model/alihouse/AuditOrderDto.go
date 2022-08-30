package alihouse

// AuditOrderDto 结构体
type AuditOrderDto struct {
	// 同意备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 统一的订单号
	HouseOrderNo string `json:"house_order_no,omitempty" xml:"house_order_no,omitempty"`
	// 同意
	IsAgree int64 `json:"is_agree,omitempty" xml:"is_agree,omitempty"`
	// 履约单id
	TradeOrderId int64 `json:"trade_order_id,omitempty" xml:"trade_order_id,omitempty"`
	// 门店id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}
