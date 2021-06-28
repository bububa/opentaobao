package wdk

// ApplyReverseResponse 
type ApplyReverseResponse struct {

    // wdk交易单号
    
    BizOrderIds   []int64 `json:"biz_order_ids,omitempty" xml:"biz_order_ids>int64,omitempty"`
    

    // 礼品卡号
    
    GiftCardNos   []string `json:"gift_card_nos,omitempty" xml:"gift_card_nos>string,omitempty"`
    

    // 1售中 2售后
    
    InSaleRefund   int64 `json:"in_sale_refund,omitempty" xml:"in_sale_refund,omitempty"`
    

    // 最大可退金额
    
    MaxRefundFee   int64 `json:"max_refund_fee,omitempty" xml:"max_refund_fee,omitempty"`
    

    // 运费
    
    PostFee   int64 `json:"post_fee,omitempty" xml:"post_fee,omitempty"`
    

    // 原因
    
    ReasonList   []ReasonVo `json:"reason_list,omitempty" xml:"reason_list,omitempty"`
    

    // 退款渠道
    
    RefundChannelList   []RefundChannelVo `json:"refund_channel_list,omitempty" xml:"refund_channel_list,omitempty"`
    

    // 请求id（提交申请入参）
    
    RequestId   string `json:"request_id,omitempty" xml:"request_id,omitempty"`
    

    // 逆向单id
    
    ReverseIds   []int64 `json:"reverse_ids,omitempty" xml:"reverse_ids>int64,omitempty"`
    

    // 门店id
    
    StoreId   string `json:"store_id,omitempty" xml:"store_id,omitempty"`
    

    // 是否支持修改金额
    
    SupportModifyAmount   bool `json:"support_modify_amount,omitempty" xml:"support_modify_amount,omitempty"`
    

}
