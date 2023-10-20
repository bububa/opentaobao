package fenxiao

import (
	"sync"
)

// ProductImageList 结构体
type ProductImageList struct {
	// 图片对应的url
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	// 当图片类型为属性图片时，表示图片对应的属性pv对。
	Properties string `json:"properties,omitempty" xml:"properties,omitempty"`
	// 图片类型（PRODUCT：产品图片  EXTPRODUCT：产品辅图  PROPERTIES：产品属性图片 ）
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 图片id
	ImageId int64 `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// 图片顺序
	ImagePosition int64 `json:"image_position,omitempty" xml:"image_position,omitempty"`
}

var poolProductImageList = sync.Pool{
	New: func() any {
		return new(ProductImageList)
	},
}

// GetProductImageList() 从对象池中获取ProductImageList
func GetProductImageList() *ProductImageList {
	return poolProductImageList.Get().(*ProductImageList)
}

// ReleaseProductImageList 释放ProductImageList
func ReleaseProductImageList(v *ProductImageList) {
	v.ImageUrl = ""
	v.Properties = ""
	v.Type = ""
	v.ImageId = 0
	v.ImagePosition = 0
	poolProductImageList.Put(v)
}
