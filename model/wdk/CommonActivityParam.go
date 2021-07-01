package wdk

// CommonActivityParam 结构体
type CommonActivityParam struct {
	// 外部活动编码
	OutActId string `json:"out_act_id,omitempty" xml:"out_act_id,omitempty"`
	// 活动Id
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
}
