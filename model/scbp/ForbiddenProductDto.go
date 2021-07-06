package scbp

// ForbiddenProductDto 结构体
type ForbiddenProductDto struct {
	// 标题
	Subject string `json:"subject,omitempty" xml:"subject,omitempty"`
	// 图片地址
	ImgUrl string `json:"img_url,omitempty" xml:"img_url,omitempty"`
	// 产品id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
