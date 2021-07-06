package icbu

// ListTemplateAPIResult 结构体
type ListTemplateAPIResult struct {
	// 运费模板集合
	Items []ShippinglineTemplate `json:"items,omitempty" xml:"items>shippingline_template,omitempty"`
	// 运费模板总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
}
