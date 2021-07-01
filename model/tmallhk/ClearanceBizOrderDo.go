package tmallhk

// ClearanceBizOrderDo 结构体
type ClearanceBizOrderDo struct {
	// 淘系订单id
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 淘系买家id
	BuyerId int64 `json:"buyer_id,omitempty" xml:"buyer_id,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 子订单列表封装
	OrderLineList []ClearanceOrderLineDo `json:"order_line_list,omitempty" xml:"order_line_list>clearance_order_line_do,omitempty"`
	// 付款状态
	PayStatus int64 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// 邮费
	PostFee int64 `json:"post_fee,omitempty" xml:"post_fee,omitempty"`
	// 退款状态
	RefundStatus int64 `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 卖家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 卖家昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 卖家旺旺
	SellerWangWangId string `json:"seller_wang_wang_id,omitempty" xml:"seller_wang_wang_id,omitempty"`
	// 税费封装
	TaxDO *ClearanceTaxDo `json:"tax_d_o,omitempty" xml:"tax_d_o,omitempty"`
	// 买家实付款
	Tf int64 `json:"tf,omitempty" xml:"tf,omitempty"`
}
