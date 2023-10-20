package feedflow

// TaobaoFeedflowItemAdgroupAdzonePageResultDto 结构体
type TaobaoFeedflowItemAdgroupAdzonePageResultDto struct {
	// 广告位列表
	AdzoneBindList []AdzoneBindDto `json:"adzone_bind_list,omitempty" xml:"adzone_bind_list>adzone_bind_dto,omitempty"`
	// 返回消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 广告位总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 调用是否成功,true-成功，false-失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
