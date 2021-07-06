package wdk

// CommonActivityRequest 结构体
type CommonActivityRequest struct {
	// 商家活动id
	OutActId string `json:"out_act_id,omitempty" xml:"out_act_id,omitempty"`
	// 五道口活动id
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
}
