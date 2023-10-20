package alitripmerchant

import (
	"sync"
)

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

var poolVoucherParameter = sync.Pool{
	New: func() any {
		return new(VoucherParameter)
	},
}

// GetVoucherParameter() 从对象池中获取VoucherParameter
func GetVoucherParameter() *VoucherParameter {
	return poolVoucherParameter.Get().(*VoucherParameter)
}

// ReleaseVoucherParameter 释放VoucherParameter
func ReleaseVoucherParameter(v *VoucherParameter) {
	v.Voucher = v.Voucher[:0]
	v.SendCouponParams = v.SendCouponParams[:0]
	v.Sign = ""
	v.SendCouponMerchant = ""
	poolVoucherParameter.Put(v)
}
