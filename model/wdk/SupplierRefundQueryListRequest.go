package wdk

// SupplierRefundQueryListRequest 
type SupplierRefundQueryListRequest struct {

    // 渠道来源
    OrderFrom   int64 `json:"order_from,omitempty"`

    // 售卖商场code
    SourceMerchantCode   string `json:"source_merchant_code,omitempty"`

    // 盒马主订单id
    MainBizOrderIds   []Number `json:"main_biz_order_ids,omitempty"`

    // 盒马子订单id
    SubBizOrderIds   []Number `json:"sub_biz_order_ids,omitempty"`

    // 退款单id
    RefundIds   []Number `json:"refund_ids,omitempty"`

}
