package promotion

import (
	"sync"
)

// TmallPromotionCouponQueryData 结构体
type TmallPromotionCouponQueryData struct {
	// discount
	Discount string `json:"discount,omitempty" xml:"discount,omitempty"`
	// startFee
	StartFee string `json:"start_fee,omitempty" xml:"start_fee,omitempty"`
	// couponName
	CouponName string `json:"coupon_name,omitempty" xml:"coupon_name,omitempty"`
	// endTime
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// startTime
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// supplierId
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// couponTemplateId
	CouponTemplateId string `json:"coupon_template_id,omitempty" xml:"coupon_template_id,omitempty"`
	// id
	Id string `json:"id,omitempty" xml:"id,omitempty"`
}

var poolTmallPromotionCouponQueryData = sync.Pool{
	New: func() any {
		return new(TmallPromotionCouponQueryData)
	},
}

// GetTmallPromotionCouponQueryData() 从对象池中获取TmallPromotionCouponQueryData
func GetTmallPromotionCouponQueryData() *TmallPromotionCouponQueryData {
	return poolTmallPromotionCouponQueryData.Get().(*TmallPromotionCouponQueryData)
}

// ReleaseTmallPromotionCouponQueryData 释放TmallPromotionCouponQueryData
func ReleaseTmallPromotionCouponQueryData(v *TmallPromotionCouponQueryData) {
	v.Discount = ""
	v.StartFee = ""
	v.CouponName = ""
	v.EndTime = ""
	v.StartTime = ""
	v.SupplierId = ""
	v.CouponTemplateId = ""
	v.Id = ""
	poolTmallPromotionCouponQueryData.Put(v)
}
