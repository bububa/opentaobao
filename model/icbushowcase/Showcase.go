package icbushowcase

// Showcase 结构体
type Showcase struct {
	// 产品描述
	Subject string `json:"subject,omitempty" xml:"subject,omitempty"`
	// 产品主图
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	// 橱窗id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 产品id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// valid
	Valid bool `json:"valid,omitempty" xml:"valid,omitempty"`
}
