package trade

// FastBuyPosCreateRequest 
type FastBuyPosCreateRequest struct {
    // 收银员工号
    CashierNum   string `json:"cashier_num,omitempty" xml:"cashier_num,omitempty"`
    // 购买商品信息
    Items   []FastBuyPosItemBo `json:"items,omitempty" xml:"items>fast_buy_pos_item_bo,omitempty"`
    // pos机id
    MachineId   string `json:"machine_id,omitempty" xml:"machine_id,omitempty"`
    // 外部唯一订单号
    OutOrderId   string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
    // 付款码
    PayCode   string `json:"pay_code,omitempty" xml:"pay_code,omitempty"`
    // 经营店id
    StoreId   string `json:"store_id,omitempty" xml:"store_id,omitempty"`
    // 支付宝支付金额
    PayFee   int64 `json:"pay_fee,omitempty" xml:"pay_fee,omitempty"`
    // 外部门店编码
    OutShopCode   string `json:"out_shop_code,omitempty" xml:"out_shop_code,omitempty"`
    // 外部优惠券支付金额
    OutCouponPayFee   int64 `json:"out_coupon_pay_fee,omitempty" xml:"out_coupon_pay_fee,omitempty"`
    // 扩展属性
    ExtendInfo   string `json:"extend_info,omitempty" xml:"extend_info,omitempty"`
    // 外部优惠金额
    OutPromotionFee   int64 `json:"out_promotion_fee,omitempty" xml:"out_promotion_fee,omitempty"`
}
