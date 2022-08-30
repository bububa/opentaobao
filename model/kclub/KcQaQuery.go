package kclub

// KcQaQuery 结构体
type KcQaQuery struct {
	// context列表
	ContextList []int64 `json:"context_list,omitempty" xml:"context_list>int64,omitempty"`
	// 问题类型列表
	QuestionTypes []string `json:"question_types,omitempty" xml:"question_types>string,omitempty"`
	// 状态列表
	StatusList []string `json:"status_list,omitempty" xml:"status_list>string,omitempty"`
	// 视角过滤
	Views []string `json:"views,omitempty" xml:"views>string,omitempty"`
	// 租户id
	TenantId int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// context
	Context int64 `json:"context,omitempty" xml:"context,omitempty"`
	// 页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 类目id
	CatId int64 `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 问题类型
	QuestionType int64 `json:"question_type,omitempty" xml:"question_type,omitempty"`
	// 排序对象
	SorterConfig *SorterConfig `json:"sorter_config,omitempty" xml:"sorter_config,omitempty"`
}
