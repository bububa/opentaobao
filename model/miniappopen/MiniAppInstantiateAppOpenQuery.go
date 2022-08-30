package miniappopen

// MiniAppInstantiateAppOpenQuery 结构体
type MiniAppInstantiateAppOpenQuery struct {
	// 模版id
	TemplateId string `json:"template_id,omitempty" xml:"template_id,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 分页
	PageNum int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
}
