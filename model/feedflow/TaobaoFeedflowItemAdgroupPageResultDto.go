package feedflow

// TaobaofeedflowitemadgrouppageResultDto 结构体
type TaobaofeedflowitemadgrouppageResultDto struct {
	// 返回数据结果
	Results []AdgroupDto `json:"results,omitempty" xml:"results>adgroup_dto,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
