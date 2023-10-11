package logistic

// TmsPictureInfoRequest 结构体
type TmsPictureInfoRequest struct {
	// 图片访问url
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// 图片上传时间
	PicUploadTime string `json:"pic_upload_time,omitempty" xml:"pic_upload_time,omitempty"`
}
