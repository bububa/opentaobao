package campus

// CalenderTemplateQuery 结构体
type CalenderTemplateQuery struct {
	// 日历模版id集合
	TemplateIdList []int64 `json:"template_id_list,omitempty" xml:"template_id_list>int64,omitempty"`
	// 序列号
	SnNo string `json:"sn_no,omitempty" xml:"sn_no,omitempty"`
}
