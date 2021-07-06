package koubeimall

// ItemImageDetail 结构体
type ItemImageDetail struct {
	// 图片列表
	ItemImages []string `json:"item_images,omitempty" xml:"item_images>string,omitempty"`
	// 描述
	ImageDescribe string `json:"image_describe,omitempty" xml:"image_describe,omitempty"`
	// 标题
	ImageTitle string `json:"image_title,omitempty" xml:"image_title,omitempty"`
}
