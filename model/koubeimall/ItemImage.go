package koubeimall

// ItemImage 结构体
type ItemImage struct {
	// 图片名称
	ImageName string `json:"image_name,omitempty" xml:"image_name,omitempty"`
	// 图片链接
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
}
