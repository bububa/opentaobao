package wdk

// CreateReverseReq 
/* model for simplify = false
type CreateReverseReq struct {

    // wdk子单号
    
    BizOrderIds  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"biz_order_ids,omitempty"`
    

    // 礼品卡号
    
    GiftCardNos  struct {
        String  []string `json:"string,omitempty"`
    } `json:"gift_card_nos,omitempty"`
    

    // 操作人
    
    Operator  *struct {
        OperatorVo  *OperatorVo `json:"operator_vo,omitempty"`
    } `json:"operator,omitempty"`
    

    // 退款凭证
    
    Proofs  struct {
        String  []string `json:"string,omitempty"`
    } `json:"proofs,omitempty"`
    

    // 退款原因id
    
    ReasonId   int64 `json:"reason_id,omitempty"`
    

    // 退款原因描述
    
    ReasonText   string `json:"reason_text,omitempty"`
    

    // 退款金额
    
    RefundAmount   int64 `json:"refund_amount,omitempty"`
    

    // 退款渠道
    
    RefundChannelList  struct {
        RefundChannelVo  []RefundChannelVo `json:"refund_channel_vo,omitempty"`
    } `json:"refund_channel_list,omitempty"`
    

    // 请求id（apply接口返回）
    
    RequestId   string `json:"request_id,omitempty"`
    

    // 门店ID
    
    StoreId   string `json:"store_id,omitempty"`
    

}
*/

// CreateReverseReq 
type CreateReverseReq struct {

    // wdk子单号
    BizOrderIds   []int64 `json:"biz_order_ids,omitempty"`

    // 礼品卡号
    GiftCardNos   []string `json:"gift_card_nos,omitempty"`

    // 操作人
    Operator   *OperatorVo `json:"operator,omitempty"`

    // 退款凭证
    Proofs   []string `json:"proofs,omitempty"`

    // 退款原因id
    ReasonId   int64 `json:"reason_id,omitempty"`

    // 退款原因描述
    ReasonText   string `json:"reason_text,omitempty"`

    // 退款金额
    RefundAmount   int64 `json:"refund_amount,omitempty"`

    // 退款渠道
    RefundChannelList   []RefundChannelVo `json:"refund_channel_list,omitempty"`

    // 请求id（apply接口返回）
    RequestId   string `json:"request_id,omitempty"`

    // 门店ID
    StoreId   string `json:"store_id,omitempty"`

}
