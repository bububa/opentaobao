package hotel

import (
	"sync"
)

// CheckInCheckOutPrice 结构体
type CheckInCheckOutPrice struct {
	// 飞猪营销优惠列表
	Promotions []PromotionPrice `json:"promotions,omitempty" xml:"promotions>promotion_price,omitempty"`
	// 入住日期 yyyy-MM-dd
	CheckinDate string `json:"checkin_date,omitempty" xml:"checkin_date,omitempty"`
	// 离店日期 yyyy-MM-dd
	CheckoutDate string `json:"checkout_date,omitempty" xml:"checkout_date,omitempty"`
	// 入住天数
	StayDays int64 `json:"stay_days,omitempty" xml:"stay_days,omitempty"`
	// 商品营销后报价，精度（分）
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 早餐数，-1代表早餐数和入住人数一致
	BreakfastCount int64 `json:"breakfast_count,omitempty" xml:"breakfast_count,omitempty"`
	// 商品营销前报价，精度（分）
	OriginalPrice int64 `json:"original_price,omitempty" xml:"original_price,omitempty"`
}

var poolCheckInCheckOutPrice = sync.Pool{
	New: func() any {
		return new(CheckInCheckOutPrice)
	},
}

// GetCheckInCheckOutPrice() 从对象池中获取CheckInCheckOutPrice
func GetCheckInCheckOutPrice() *CheckInCheckOutPrice {
	return poolCheckInCheckOutPrice.Get().(*CheckInCheckOutPrice)
}

// ReleaseCheckInCheckOutPrice 释放CheckInCheckOutPrice
func ReleaseCheckInCheckOutPrice(v *CheckInCheckOutPrice) {
	v.Promotions = v.Promotions[:0]
	v.CheckinDate = ""
	v.CheckoutDate = ""
	v.StayDays = 0
	v.Price = 0
	v.BreakfastCount = 0
	v.OriginalPrice = 0
	poolCheckInCheckOutPrice.Put(v)
}
