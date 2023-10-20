package shenjing

import (
	"sync"
)

// UploadFaceDo 结构体
type UploadFaceDo struct {
	// 中文消息
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// extDesc
	ExtDesc string `json:"ext_desc,omitempty" xml:"ext_desc,omitempty"`
	// 状态码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 访客总数
	VisitorFacePhotoTotalNum int64 `json:"visitor_face_photo_total_num,omitempty" xml:"visitor_face_photo_total_num,omitempty"`
	// 出参
	Map *ResultMap `json:"map,omitempty" xml:"map,omitempty"`
	// 当前上传人数
	VisitorFacePhotoCurrentNum int64 `json:"visitor_face_photo_current_num,omitempty" xml:"visitor_face_photo_current_num,omitempty"`
}

var poolUploadFaceDo = sync.Pool{
	New: func() any {
		return new(UploadFaceDo)
	},
}

// GetUploadFaceDo() 从对象池中获取UploadFaceDo
func GetUploadFaceDo() *UploadFaceDo {
	return poolUploadFaceDo.Get().(*UploadFaceDo)
}

// ReleaseUploadFaceDo 释放UploadFaceDo
func ReleaseUploadFaceDo(v *UploadFaceDo) {
	v.Desc = ""
	v.ExtDesc = ""
	v.Code = ""
	v.VisitorFacePhotoTotalNum = 0
	v.Map = nil
	v.VisitorFacePhotoCurrentNum = 0
	poolUploadFaceDo.Put(v)
}
