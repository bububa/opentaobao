package koubeimall

import (
	"sync"
)

// ItemImageDetail 结构体
type ItemImageDetail struct {
	// 图片列表
	ItemImages []string `json:"item_images,omitempty" xml:"item_images>string,omitempty"`
	// 描述
	ImageDescribe string `json:"image_describe,omitempty" xml:"image_describe,omitempty"`
	// 标题
	ImageTitle string `json:"image_title,omitempty" xml:"image_title,omitempty"`
}

var poolItemImageDetail = sync.Pool{
	New: func() any {
		return new(ItemImageDetail)
	},
}

// GetItemImageDetail() 从对象池中获取ItemImageDetail
func GetItemImageDetail() *ItemImageDetail {
	return poolItemImageDetail.Get().(*ItemImageDetail)
}

// ReleaseItemImageDetail 释放ItemImageDetail
func ReleaseItemImageDetail(v *ItemImageDetail) {
	v.ItemImages = v.ItemImages[:0]
	v.ImageDescribe = ""
	v.ImageTitle = ""
	poolItemImageDetail.Put(v)
}
