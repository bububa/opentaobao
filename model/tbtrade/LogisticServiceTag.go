package tbtrade

// LogisticServiceTag 结构体
type LogisticServiceTag struct {
	// 物流服务下的标签属性,多个标签之间有&#34;;&#34;分隔
	ServiceTag string `json:"service_tag,omitempty" xml:"service_tag,omitempty"`
	// 消费者选快递请直接判断service_tag是否包含companyCode。而不要判断service_type
	ServiceType string `json:"service_type,omitempty" xml:"service_type,omitempty"`
}
