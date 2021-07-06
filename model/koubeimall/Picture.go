package koubeimall

// Picture 结构体
type Picture struct {
	// 图片来源
	PictureSource string `json:"picture_source,omitempty" xml:"picture_source,omitempty"`
	// 图片链接
	PictureUrl string `json:"picture_url,omitempty" xml:"picture_url,omitempty"`
	// 图片名称
	PictureName string `json:"picture_name,omitempty" xml:"picture_name,omitempty"`
	// 图片顺序
	Sequence int64 `json:"sequence,omitempty" xml:"sequence,omitempty"`
}
