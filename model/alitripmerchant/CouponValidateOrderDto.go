package alitripmerchant

// CouponValidateOrderDto 结构体
type CouponValidateOrderDto struct {
	// 优惠券实例列表
	InstanceVOList []CouponInfoVo `json:"instance_v_o_list,omitempty" xml:"instance_v_o_list>coupon_info_vo,omitempty"`
	// 优惠券实例对象
	CouponInfoVO *CouponInfoVo `json:"coupon_info_v_o,omitempty" xml:"coupon_info_v_o,omitempty"`
	// 折扣价试单结果
	DiscountDTO *ValidateResultVo `json:"discount_d_t_o,omitempty" xml:"discount_d_t_o,omitempty"`
	// 原价试单结果
	NoDiscountDTO *ValidateResultVo `json:"no_discount_d_t_o,omitempty" xml:"no_discount_d_t_o,omitempty"`
	// 权益商品展示信息
	DerbyVoucherInfo *DerbyVoucherInfo `json:"derby_voucher_info,omitempty" xml:"derby_voucher_info,omitempty"`
}
