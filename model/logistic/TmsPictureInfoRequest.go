package logistic

import (
	"sync"
)

// TmsPictureInfoRequest 结构体
type TmsPictureInfoRequest struct {
	// 图片访问url
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// 图片上传时间
	PicUploadTime string `json:"pic_upload_time,omitempty" xml:"pic_upload_time,omitempty"`
}

var poolTmsPictureInfoRequest = sync.Pool{
	New: func() any {
		return new(TmsPictureInfoRequest)
	},
}

// GetTmsPictureInfoRequest() 从对象池中获取TmsPictureInfoRequest
func GetTmsPictureInfoRequest() *TmsPictureInfoRequest {
	return poolTmsPictureInfoRequest.Get().(*TmsPictureInfoRequest)
}

// ReleaseTmsPictureInfoRequest 释放TmsPictureInfoRequest
func ReleaseTmsPictureInfoRequest(v *TmsPictureInfoRequest) {
	v.PicUrl = ""
	v.PicUploadTime = ""
	poolTmsPictureInfoRequest.Put(v)
}
