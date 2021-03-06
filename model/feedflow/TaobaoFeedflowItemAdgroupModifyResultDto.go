package feedflow

// TaobaoFeedflowItemAdgroupModifyResultDto 结构体
type TaobaoFeedflowItemAdgroupModifyResultDto struct {
	// 信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 单元id
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
