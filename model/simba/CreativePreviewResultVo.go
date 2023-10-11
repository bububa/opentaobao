package simba

// CreativePreviewResultVo 结构体
type CreativePreviewResultVo struct {
	// 预览创意图片地址
	PreviewImgUrl string `json:"preview_img_url,omitempty" xml:"preview_img_url,omitempty"`
	// 预览落地页地址
	PreviewClickUrl string `json:"preview_click_url,omitempty" xml:"preview_click_url,omitempty"`
	// 预览创意尺寸
	PreviewSize string `json:"preview_size,omitempty" xml:"preview_size,omitempty"`
	// 预览创意尺寸
	PreviewScale string `json:"preview_scale,omitempty" xml:"preview_scale,omitempty"`
	// 视频地址
	PreviewVideoPath string `json:"preview_video_path,omitempty" xml:"preview_video_path,omitempty"`
	// 视频封面
	PreviewVideoImage string `json:"preview_video_image,omitempty" xml:"preview_video_image,omitempty"`
	// 预览创意类型,2:图片,10:创意模板,12:微视频
	PreviewFormat int64 `json:"preview_format,omitempty" xml:"preview_format,omitempty"`
}
