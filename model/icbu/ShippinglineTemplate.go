package icbu

// ShippinglineTemplate 结构体
type ShippinglineTemplate struct {
	// 运费模板id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 运费模板名称
	Title string `json:"title,omitempty" xml:"title,omitempty"`
}
