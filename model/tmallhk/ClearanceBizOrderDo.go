package tmallhk

// ClearanceBizOrderDo 
type ClearanceBizOrderDo struct {

    // 淘系订单id
    BizOrderId   int64 `json:"biz_order_id,omitempty"`

    // 淘系买家id
    BuyerId   int64 `json:"buyer_id,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 子订单列表封装
    OrderLineList   []ClearanceOrderLineDo `json:"order_line_list,omitempty"`

    // 付款状态
    PayStatus   int64 `json:"pay_status,omitempty"`

    // 邮费
    PostFee   int64 `json:"post_fee,omitempty"`

    // 退款状态
    RefundStatus   int64 `json:"refund_status,omitempty"`

    // 卖家id
    SellerId   int64 `json:"seller_id,omitempty"`

    // 卖家昵称
    SellerNick   string `json:"seller_nick,omitempty"`

    // 卖家旺旺
    SellerWangWangId   string `json:"seller_wang_wang_id,omitempty"`

    // 税费封装
    TaxDO   *ClearanceTaxDo `json:"tax_d_o,omitempty"`

    // 买家实付款
    Tf   int64 `json:"tf,omitempty"`

}
