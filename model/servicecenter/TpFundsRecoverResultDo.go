package servicecenter

// TpFundsRecoverResultDo 
type TpFundsRecoverResultDo struct {
    // 实际扣回金额，单位分
    ActualRecoverAmount   int64 `json:"actual_recover_amount,omitempty" xml:"actual_recover_amount,omitempty"`
    // 应扣回金额，单位分
    ToRecoverAmount   int64 `json:"to_recover_amount,omitempty" xml:"to_recover_amount,omitempty"`
    // 订单ID
    BizOrderId   int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}
