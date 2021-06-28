package wdk

// SupplierRefundQueryListRequest 
/* model for simplify = false
type SupplierRefundQueryListRequest struct {

    // 渠道来源
    
    OrderFrom   int64 `json:"order_from,omitempty"`
    

    // 售卖商场code
    
    SourceMerchantCode   string `json:"source_merchant_code,omitempty"`
    

    // 盒马主订单id
    
    MainBizOrderIds  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"main_biz_order_ids,omitempty"`
    

    // 盒马子订单id
    
    SubBizOrderIds  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"sub_biz_order_ids,omitempty"`
    

    // 退款单id
    
    RefundIds  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"refund_ids,omitempty"`
    

}
*/

// SupplierRefundQueryListRequest 
type SupplierRefundQueryListRequest struct {

    // 渠道来源
    OrderFrom   int64 `json:"order_from,omitempty"`

    // 售卖商场code
    SourceMerchantCode   string `json:"source_merchant_code,omitempty"`

    // 盒马主订单id
    MainBizOrderIds   []int64 `json:"main_biz_order_ids,omitempty"`

    // 盒马子订单id
    SubBizOrderIds   []int64 `json:"sub_biz_order_ids,omitempty"`

    // 退款单id
    RefundIds   []int64 `json:"refund_ids,omitempty"`

}
