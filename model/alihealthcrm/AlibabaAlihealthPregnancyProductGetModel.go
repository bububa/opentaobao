package alihealthcrm

// AlibabaAlihealthPregnancyProductGetModel 结构体
type AlibabaAlihealthPregnancyProductGetModel struct {
	// 文章列表
	Contents []Contents `json:"contents,omitempty" xml:"contents>contents,omitempty"`
	// 总页数
	TotalPage string `json:"total_page,omitempty" xml:"total_page,omitempty"`
	// 总条数
	TotalCount string `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
