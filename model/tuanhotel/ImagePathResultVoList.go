package tuanhotel

import (
	"sync"
)

// ImagePathResultVoList 结构体
type ImagePathResultVoList struct {
	// 图片id
	ImageUid string `json:"image_uid,omitempty" xml:"image_uid,omitempty"`
	// 在图片中心的地址路径
	ImagePath string `json:"image_path,omitempty" xml:"image_path,omitempty"`
	// 异常信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}

var poolImagePathResultVoList = sync.Pool{
	New: func() any {
		return new(ImagePathResultVoList)
	},
}

// GetImagePathResultVoList() 从对象池中获取ImagePathResultVoList
func GetImagePathResultVoList() *ImagePathResultVoList {
	return poolImagePathResultVoList.Get().(*ImagePathResultVoList)
}

// ReleaseImagePathResultVoList 释放ImagePathResultVoList
func ReleaseImagePathResultVoList(v *ImagePathResultVoList) {
	v.ImageUid = ""
	v.ImagePath = ""
	v.ErrorMsg = ""
	poolImagePathResultVoList.Put(v)
}
