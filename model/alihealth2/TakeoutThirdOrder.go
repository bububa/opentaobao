package alihealth2

// TakeoutThirdOrder 
type TakeoutThirdOrder struct {
    // 实际支付金额
    ActualPayFee   int64 `json:"actual_pay_fee,omitempty" xml:"actual_pay_fee,omitempty"`
    // 店铺id
    StoreId   int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
    // 支付时间
    PayTime   string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
    // 订单失败原因
    FaildReason   string `json:"faild_reason,omitempty" xml:"faild_reason,omitempty"`
    // 订单商品
    GoodsList   []GoodsList `json:"goods_list,omitempty" xml:"goods_list>goods_list,omitempty"`
}
