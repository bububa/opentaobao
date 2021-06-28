package alsc

// BackflowRequest 
type BackflowRequest struct {

    // 订单来源 KERUYUN  CHENSEN  KPOS
    
    BizSource   string `json:"biz_source,omitempty" xml:"biz_source,omitempty"`
    

    // 订单或子订单属性信息
    
    OrderAttributeInfoList   []OrderAttributeInfo `json:"order_attribute_info_list,omitempty" xml:"order_attribute_info_list,omitempty"`
    

    // 订单信息
    
    OrderInfo   *OrderInfo `json:"order_info,omitempty" xml:"order_info,omitempty"`
    

    // 支付流水明细
    
    PayDetailInfoList   []PayDetailInfo `json:"pay_detail_info_list,omitempty" xml:"pay_detail_info_list,omitempty"`
    

    // 优惠明细
    
    PromoDetailInfoList   []PromoDetailInfo `json:"promo_detail_info_list,omitempty" xml:"promo_detail_info_list,omitempty"`
    

    // 退款资金流水
    
    RefundFundDetailInfoList   []RefundFundDetailInfo `json:"refund_fund_detail_info_list,omitempty" xml:"refund_fund_detail_info_list,omitempty"`
    

    // 退款商品明细
    
    RefundItemDetailInfoList   []RefundItemDetailInfo `json:"refund_item_detail_info_list,omitempty" xml:"refund_item_detail_info_list,omitempty"`
    

    // 退款单数据
    
    RefundOrderInfoList   []RefundOrderInfo `json:"refund_order_info_list,omitempty" xml:"refund_order_info_list,omitempty"`
    

    // 子订单（商品）信息
    
    SubOrderInfoList   []SubOrderInfo `json:"sub_order_info_list,omitempty" xml:"sub_order_info_list,omitempty"`
    

}
