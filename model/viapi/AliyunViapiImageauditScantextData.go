package viapi

// AliyunViapiImageauditScantextData 结构体
type AliyunViapiImageauditScantextData struct {
	// 检测结果各个子元素
	Elements []Element `json:"elements,omitempty" xml:"elements>element,omitempty"`
}
