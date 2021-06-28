package alsc

// BackflowRequest 
/* model for simplify = false
type BackflowRequest struct {

    // 订单来源 KERUYUN  CHENSEN  KPOS
    
    BizSource   string `json:"biz_source,omitempty"`
    

    // 订单或子订单属性信息
    
    OrderAttributeInfoList  struct {
        OrderAttributeInfo  []OrderAttributeInfo `json:"order_attribute_info,omitempty"`
    } `json:"order_attribute_info_list,omitempty"`
    

    // 订单信息
    
    OrderInfo  *struct {
        OrderInfo  *OrderInfo `json:"order_info,omitempty"`
    } `json:"order_info,omitempty"`
    

    // 支付流水明细
    
    PayDetailInfoList  struct {
        PayDetailInfo  []PayDetailInfo `json:"pay_detail_info,omitempty"`
    } `json:"pay_detail_info_list,omitempty"`
    

    // 优惠明细
    
    PromoDetailInfoList  struct {
        PromoDetailInfo  []PromoDetailInfo `json:"promo_detail_info,omitempty"`
    } `json:"promo_detail_info_list,omitempty"`
    

    // 退款资金流水
    
    RefundFundDetailInfoList  struct {
        RefundFundDetailInfo  []RefundFundDetailInfo `json:"refund_fund_detail_info,omitempty"`
    } `json:"refund_fund_detail_info_list,omitempty"`
    

    // 退款商品明细
    
    RefundItemDetailInfoList  struct {
        RefundItemDetailInfo  []RefundItemDetailInfo `json:"refund_item_detail_info,omitempty"`
    } `json:"refund_item_detail_info_list,omitempty"`
    

    // 退款单数据
    
    RefundOrderInfoList  struct {
        RefundOrderInfo  []RefundOrderInfo `json:"refund_order_info,omitempty"`
    } `json:"refund_order_info_list,omitempty"`
    

    // 子订单（商品）信息
    
    SubOrderInfoList  struct {
        SubOrderInfo  []SubOrderInfo `json:"sub_order_info,omitempty"`
    } `json:"sub_order_info_list,omitempty"`
    

}
*/

// BackflowRequest 
type BackflowRequest struct {

    // 订单来源 KERUYUN  CHENSEN  KPOS
    BizSource   string `json:"biz_source,omitempty"`

    // 订单或子订单属性信息
    OrderAttributeInfoList   []OrderAttributeInfo `json:"order_attribute_info_list,omitempty"`

    // 订单信息
    OrderInfo   *OrderInfo `json:"order_info,omitempty"`

    // 支付流水明细
    PayDetailInfoList   []PayDetailInfo `json:"pay_detail_info_list,omitempty"`

    // 优惠明细
    PromoDetailInfoList   []PromoDetailInfo `json:"promo_detail_info_list,omitempty"`

    // 退款资金流水
    RefundFundDetailInfoList   []RefundFundDetailInfo `json:"refund_fund_detail_info_list,omitempty"`

    // 退款商品明细
    RefundItemDetailInfoList   []RefundItemDetailInfo `json:"refund_item_detail_info_list,omitempty"`

    // 退款单数据
    RefundOrderInfoList   []RefundOrderInfo `json:"refund_order_info_list,omitempty"`

    // 子订单（商品）信息
    SubOrderInfoList   []SubOrderInfo `json:"sub_order_info_list,omitempty"`

}
