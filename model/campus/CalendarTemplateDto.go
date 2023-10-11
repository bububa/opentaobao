package campus

// CalendarTemplateDto 结构体
type CalendarTemplateDto struct {
	// 时期集合
	DateList []string `json:"date_list,omitempty" xml:"date_list>string,omitempty"`
	// 日期名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 日期ID
	TemplateId int64 `json:"template_id,omitempty" xml:"template_id,omitempty"`
}
