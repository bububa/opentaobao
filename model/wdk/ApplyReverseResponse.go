package wdk

// ApplyReverseResponse 
/* model for simplify = false
type ApplyReverseResponse struct {

    // wdk交易单号
    
    BizOrderIds  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"biz_order_ids,omitempty"`
    

    // 礼品卡号
    
    GiftCardNos  struct {
        String  []string `json:"string,omitempty"`
    } `json:"gift_card_nos,omitempty"`
    

    // 1售中 2售后
    
    InSaleRefund   int64 `json:"in_sale_refund,omitempty"`
    

    // 最大可退金额
    
    MaxRefundFee   int64 `json:"max_refund_fee,omitempty"`
    

    // 运费
    
    PostFee   int64 `json:"post_fee,omitempty"`
    

    // 原因
    
    ReasonList  struct {
        ReasonVo  []ReasonVo `json:"reason_vo,omitempty"`
    } `json:"reason_list,omitempty"`
    

    // 退款渠道
    
    RefundChannelList  struct {
        RefundChannelVo  []RefundChannelVo `json:"refund_channel_vo,omitempty"`
    } `json:"refund_channel_list,omitempty"`
    

    // 请求id（提交申请入参）
    
    RequestId   string `json:"request_id,omitempty"`
    

    // 逆向单id
    
    ReverseIds  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"reverse_ids,omitempty"`
    

    // 门店id
    
    StoreId   string `json:"store_id,omitempty"`
    

    // 是否支持修改金额
    
    SupportModifyAmount   bool `json:"support_modify_amount,omitempty"`
    

}
*/

// ApplyReverseResponse 
type ApplyReverseResponse struct {

    // wdk交易单号
    BizOrderIds   []int64 `json:"biz_order_ids,omitempty"`

    // 礼品卡号
    GiftCardNos   []string `json:"gift_card_nos,omitempty"`

    // 1售中 2售后
    InSaleRefund   int64 `json:"in_sale_refund,omitempty"`

    // 最大可退金额
    MaxRefundFee   int64 `json:"max_refund_fee,omitempty"`

    // 运费
    PostFee   int64 `json:"post_fee,omitempty"`

    // 原因
    ReasonList   []ReasonVo `json:"reason_list,omitempty"`

    // 退款渠道
    RefundChannelList   []RefundChannelVo `json:"refund_channel_list,omitempty"`

    // 请求id（提交申请入参）
    RequestId   string `json:"request_id,omitempty"`

    // 逆向单id
    ReverseIds   []int64 `json:"reverse_ids,omitempty"`

    // 门店id
    StoreId   string `json:"store_id,omitempty"`

    // 是否支持修改金额
    SupportModifyAmount   bool `json:"support_modify_amount,omitempty"`

}
