package tmallnr

// NrCancelFulfillReqDto 
type NrCancelFulfillReqDto struct {
    // 操作取消人员姓名
    CancelOperUserName   string `json:"cancel_oper_user_name,omitempty" xml:"cancel_oper_user_name,omitempty"`
    // 操作取消人员ID号
    CancelOperUserId   int64 `json:"cancel_oper_user_id,omitempty" xml:"cancel_oper_user_id,omitempty"`
    // 取消的对应编码
    CancelReasonCode   int64 `json:"cancel_reason_code,omitempty" xml:"cancel_reason_code,omitempty"`
    // 淘宝商家的sellerID
    SellerId   int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
    // 取消类型。操作人类型:1寄件人,3客服小二,4快递员, 5CP, 6收件人,8门店取消,100系统
    CancelOperType   int64 `json:"cancel_oper_type,omitempty" xml:"cancel_oper_type,omitempty"`
    // 淘宝交易的主订单号
    MainOrderId   int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
    // 取消原因的说明
    CancelReason   string `json:"cancel_reason,omitempty" xml:"cancel_reason,omitempty"`
}
