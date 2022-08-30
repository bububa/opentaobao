package wlb

// CustomData 结构体
type CustomData struct {
	//  自定义区数据
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	//  自定义区链接
	TemplateUrl string `json:"template_url,omitempty" xml:"template_url,omitempty"`
}
