package feedflow

// TaobaoFeedflowItemAdgroupDeleteResultDto 结构体
type TaobaoFeedflowItemAdgroupDeleteResultDto struct {
	// 信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 删除结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
