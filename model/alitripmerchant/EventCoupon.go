package alitripmerchant

import (
	"sync"
)

// EventCoupon 结构体
type EventCoupon struct {
	// true
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// 1
	HomePagePopup string `json:"home_page_popup,omitempty" xml:"home_page_popup,omitempty"`
	// 1
	DetailsPagePicture string `json:"details_page_picture,omitempty" xml:"details_page_picture,omitempty"`
	// 1
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// true
	WhetherToIssueCoupons bool `json:"whether_to_issue_coupons,omitempty" xml:"whether_to_issue_coupons,omitempty"`
}

var poolEventCoupon = sync.Pool{
	New: func() any {
		return new(EventCoupon)
	},
}

// GetEventCoupon() 从对象池中获取EventCoupon
func GetEventCoupon() *EventCoupon {
	return poolEventCoupon.Get().(*EventCoupon)
}

// ReleaseEventCoupon 释放EventCoupon
func ReleaseEventCoupon(v *EventCoupon) {
	v.ActivityName = ""
	v.HomePagePopup = ""
	v.DetailsPagePicture = ""
	v.ActivityId = 0
	v.WhetherToIssueCoupons = false
	poolEventCoupon.Put(v)
}
