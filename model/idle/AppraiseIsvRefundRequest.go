package idle

// AppraiseIsvRefundRequest 
type AppraiseIsvRefundRequest struct {
    // 订单号
    BizOrderId   int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
    // 退款操作说明
    LeaveMessage   string `json:"leave_message,omitempty" xml:"leave_message,omitempty"`
    // 退款操作/**      * 同意退款并执行打款->退款成功，仅退款情况下，直接调用；退货退款情况下，先调用同意退货接口，收到货再调用这个接口      */     AGREE_REFUND,     /**      * 拒绝退款 -> 退款关闭      */     REFUSE_REFUND,     /**      * 同意退货并提供退货地址 -> 等待买家退货      */     AGREE_RETURN
    Operation   string `json:"operation,omitempty" xml:"operation,omitempty"`
    // 退货地址
    ReturnGoodsAddress   *ShippingAddressInfo `json:"return_goods_address,omitempty" xml:"return_goods_address,omitempty"`
}
