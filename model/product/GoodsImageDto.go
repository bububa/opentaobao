package product

// GoodsImageDto 结构体
type GoodsImageDto struct {
	// 原图url
	OriginImage string `json:"origin_image,omitempty" xml:"origin_image,omitempty"`
	// 备注
	Note string `json:"note,omitempty" xml:"note,omitempty"`
	// 图片id
	ImageId int64 `json:"image_id,omitempty" xml:"image_id,omitempty"`
}
