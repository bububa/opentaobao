package usergrowth

// RecognitionParam 结构体
type RecognitionParam struct {
	// 图片url
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// 媒体站点id
	SiteId string `json:"site_id,omitempty" xml:"site_id,omitempty"`
	// 图片id
	PicId string `json:"pic_id,omitempty" xml:"pic_id,omitempty"`
}
