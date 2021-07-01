package tuanhotel

// ImageInfoVOList 结构体
type ImageInfoVOList struct {
	// 图片ID
	ImageUid string `json:"image_uid,omitempty" xml:"image_uid,omitempty"`
	// 图片URL地址
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	// false
	IsPhone bool `json:"is_phone,omitempty" xml:"is_phone,omitempty"`
}
