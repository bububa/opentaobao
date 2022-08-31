package alitripmerchant

// VoucherParameter 结构体
type VoucherParameter struct {
	// 套餐券实体类
	Voucher []VoucherInfoVo `json:"voucher,omitempty" xml:"voucher>voucher_info_vo,omitempty"`
	// 添加优惠券参至卡包参数
	SendCouponParams []SendCouponParams `json:"send_coupon_params,omitempty" xml:"send_coupon_params>send_coupon_params,omitempty"`
	// 加密签名
	Sign string `json:"sign,omitempty" xml:"sign,omitempty"`
	// 商户号
	SendCouponMerchant string `json:"send_coupon_merchant,omitempty" xml:"send_coupon_merchant,omitempty"`
}
