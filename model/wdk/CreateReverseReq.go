package wdk

// CreateReverseReq 
type CreateReverseReq struct {

    // wdk子单号
    BizOrderIds   []Number `json:"biz_order_ids,omitempty"`

    // 礼品卡号
    GiftCardNos   []String `json:"gift_card_nos,omitempty"`

    // 操作人
    Operator   *OperatorVo `json:"operator,omitempty"`

    // 退款凭证
    Proofs   []String `json:"proofs,omitempty"`

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
