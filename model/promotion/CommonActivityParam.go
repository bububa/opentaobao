package promotion

// CommonActivityParam 结构体
type CommonActivityParam struct {
	// 五道口优惠券活动id
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 商家优惠券活动id
	OutActId string `json:"out_act_id,omitempty" xml:"out_act_id,omitempty"`
}
