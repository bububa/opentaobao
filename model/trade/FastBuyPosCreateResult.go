package trade

import (
	"sync"
)

// FastBuyPosCreateResult 结构体
type FastBuyPosCreateResult struct {
	// 返回错误码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 返回错误信息
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 盒马订单号
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 优惠券优惠金额
	CouponFee int64 `json:"coupon_fee,omitempty" xml:"coupon_fee,omitempty"`
	// 优惠活动优惠金额
	PromotionFee int64 `json:"promotion_fee,omitempty" xml:"promotion_fee,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolFastBuyPosCreateResult = sync.Pool{
	New: func() any {
		return new(FastBuyPosCreateResult)
	},
}

// GetFastBuyPosCreateResult() 从对象池中获取FastBuyPosCreateResult
func GetFastBuyPosCreateResult() *FastBuyPosCreateResult {
	return poolFastBuyPosCreateResult.Get().(*FastBuyPosCreateResult)
}

// ReleaseFastBuyPosCreateResult 释放FastBuyPosCreateResult
func ReleaseFastBuyPosCreateResult(v *FastBuyPosCreateResult) {
	v.ReturnCode = ""
	v.ReturnMsg = ""
	v.BizOrderId = 0
	v.CouponFee = 0
	v.PromotionFee = 0
	v.Success = false
	poolFastBuyPosCreateResult.Put(v)
}
