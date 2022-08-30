package wdk

// ItemPoolActivityQueryRequest 结构体
type ItemPoolActivityQueryRequest struct {
	// 外部活动id
	OutActId string `json:"out_act_id,omitempty" xml:"out_act_id,omitempty"`
	// 活动ID
	ActId int64 `json:"act_id,omitempty" xml:"act_id,omitempty"`
}
