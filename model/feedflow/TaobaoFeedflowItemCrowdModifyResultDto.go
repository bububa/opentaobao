package feedflow

// TaobaoFeedflowItemCrowdModifyResultDto 结构体
type TaobaoFeedflowItemCrowdModifyResultDto struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 调用是否成功,true-成功，false-失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
