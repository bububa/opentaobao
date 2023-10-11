package tmallhk

// ClearanceLogisticsOrderDo 结构体
type ClearanceLogisticsOrderDo struct {
	// 子订单列表封装
	OrderLineList []ClearanceOrderLineDo `json:"order_line_list,omitempty" xml:"order_line_list>clearance_order_line_do,omitempty"`
	// 清关订单编号
	ClearanceOrderNo string `json:"clearance_order_no,omitempty" xml:"clearance_order_no,omitempty"`
	// 税费封装
	TaxDO *ClearanceTaxDo `json:"tax_d_o,omitempty" xml:"tax_d_o,omitempty"`
	// 邮费
	PostFee int64 `json:"post_fee,omitempty" xml:"post_fee,omitempty"`
	// 买家id
	BuyerId int64 `json:"buyer_id,omitempty" xml:"buyer_id,omitempty"`
	// 买家实付款
	Tf int64 `json:"tf,omitempty" xml:"tf,omitempty"`
	// 订单id
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}
