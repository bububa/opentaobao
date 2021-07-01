package ieagency

// RefundActivityVo 结构体
type RefundActivityVo struct {
	// 活动名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 活动收回金额(单位：分)
	TakeBackPrice int64 `json:"take_back_price,omitempty" xml:"take_back_price,omitempty"`
}
