package tuanhotel

import (
	"sync"
)

// ImageInfoVOList 结构体
type ImageInfoVOList struct {
	// 图片ID
	ImageUid string `json:"image_uid,omitempty" xml:"image_uid,omitempty"`
	// 图片URL地址
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	// false
	IsPhone bool `json:"is_phone,omitempty" xml:"is_phone,omitempty"`
}

var poolImageInfoVOList = sync.Pool{
	New: func() any {
		return new(ImageInfoVOList)
	},
}

// GetImageInfoVOList() 从对象池中获取ImageInfoVOList
func GetImageInfoVOList() *ImageInfoVOList {
	return poolImageInfoVOList.Get().(*ImageInfoVOList)
}

// ReleaseImageInfoVOList 释放ImageInfoVOList
func ReleaseImageInfoVOList(v *ImageInfoVOList) {
	v.ImageUid = ""
	v.ImageUrl = ""
	v.IsPhone = false
	poolImageInfoVOList.Put(v)
}
