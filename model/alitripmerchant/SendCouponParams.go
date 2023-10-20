package alitripmerchant

import (
	"sync"
)

// SendCouponParams 结构体
type SendCouponParams struct {
	// 1
	StockId string `json:"stock_id,omitempty" xml:"stock_id,omitempty"`
	// 1
	OutRequestNo string `json:"out_request_no,omitempty" xml:"out_request_no,omitempty"`
	// 1
	CouponCode string `json:"coupon_code,omitempty" xml:"coupon_code,omitempty"`
}

var poolSendCouponParams = sync.Pool{
	New: func() any {
		return new(SendCouponParams)
	},
}

// GetSendCouponParams() 从对象池中获取SendCouponParams
func GetSendCouponParams() *SendCouponParams {
	return poolSendCouponParams.Get().(*SendCouponParams)
}

// ReleaseSendCouponParams 释放SendCouponParams
func ReleaseSendCouponParams(v *SendCouponParams) {
	v.StockId = ""
	v.OutRequestNo = ""
	v.CouponCode = ""
	poolSendCouponParams.Put(v)
}
