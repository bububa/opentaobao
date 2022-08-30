package alitripmerchant

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
