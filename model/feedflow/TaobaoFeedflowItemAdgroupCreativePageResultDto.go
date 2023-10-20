package feedflow

// TaobaofeedflowitemadgroupcreativepageResultDto 结构体
type TaobaofeedflowitemadgroupcreativepageResultDto struct {
	// 绑定创意的列表
	CreativeBindList []CreativeBindDto `json:"creative_bind_list,omitempty" xml:"creative_bind_list>creative_bind_dto,omitempty"`
	// 消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 总数目
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 调用是否成功,true-成功，false-失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
