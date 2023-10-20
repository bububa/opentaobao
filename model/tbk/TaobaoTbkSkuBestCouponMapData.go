package tbk

import (
	"sync"
)

// TaobaoTbkSkuBestCouponMapData 结构体
type TaobaoTbkSkuBestCouponMapData struct {
	// 优惠券过期时间13位时间戳
	CouponExpireTime int64 `json:"coupon_expire_time,omitempty" xml:"coupon_expire_time,omitempty"`
}

var poolTaobaoTbkSkuBestCouponMapData = sync.Pool{
	New: func() any {
		return new(TaobaoTbkSkuBestCouponMapData)
	},
}

// GetTaobaoTbkSkuBestCouponMapData() 从对象池中获取TaobaoTbkSkuBestCouponMapData
func GetTaobaoTbkSkuBestCouponMapData() *TaobaoTbkSkuBestCouponMapData {
	return poolTaobaoTbkSkuBestCouponMapData.Get().(*TaobaoTbkSkuBestCouponMapData)
}

// ReleaseTaobaoTbkSkuBestCouponMapData 释放TaobaoTbkSkuBestCouponMapData
func ReleaseTaobaoTbkSkuBestCouponMapData(v *TaobaoTbkSkuBestCouponMapData) {
	v.CouponExpireTime = 0
	poolTaobaoTbkSkuBestCouponMapData.Put(v)
}
