package wdk

// SupplierRefundQueryListRequest 
type SupplierRefundQueryListRequest struct {
    // 渠道来源
    OrderFrom   int64 `json:"order_from,omitempty" xml:"order_from,omitempty"`
    // 售卖商场code
    SourceMerchantCode   string `json:"source_merchant_code,omitempty" xml:"source_merchant_code,omitempty"`
    // 盒马主订单id
    MainBizOrderIds   []int64 `json:"main_biz_order_ids,omitempty" xml:"main_biz_order_ids>int64,omitempty"`
    // 盒马子订单id
    SubBizOrderIds   []int64 `json:"sub_biz_order_ids,omitempty" xml:"sub_biz_order_ids>int64,omitempty"`
    // 退款单id
    RefundIds   []int64 `json:"refund_ids,omitempty" xml:"refund_ids>int64,omitempty"`
}
