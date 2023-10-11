package idle

// IdleAdvMaterialUploadDetailTopParam 结构体
type IdleAdvMaterialUploadDetailTopParam struct {
	// 素材地址，必传字段, 如果链接有有效期，尽量有效期放开到1个小时
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 素材名称
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 如果是视频素材，需要传当前视频的封面图地址,视频素材必填项
	CoverUrl string `json:"cover_url,omitempty" xml:"cover_url,omitempty"`
	// 素材类型，1 图片 2 视频
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 必填项,如果是视频素材，需要传当前视频的封面图的宽,如果是图片素材，需要传当前图片的宽
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
	// 必填项,如果是视频素材，需要传当前视频的封面图的宽,如果是图片素材，需要传当前图片的高
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
}
