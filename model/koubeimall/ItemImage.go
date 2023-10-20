package koubeimall

import (
	"sync"
)

// ItemImage 结构体
type ItemImage struct {
	// 图片名称
	ImageName string `json:"image_name,omitempty" xml:"image_name,omitempty"`
	// 图片链接
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
}

var poolItemImage = sync.Pool{
	New: func() any {
		return new(ItemImage)
	},
}

// GetItemImage() 从对象池中获取ItemImage
func GetItemImage() *ItemImage {
	return poolItemImage.Get().(*ItemImage)
}

// ReleaseItemImage 释放ItemImage
func ReleaseItemImage(v *ItemImage) {
	v.ImageName = ""
	v.ImageUrl = ""
	poolItemImage.Put(v)
}
