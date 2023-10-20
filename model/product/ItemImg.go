package product

import (
	"sync"
)

// ItemImg 结构体
type ItemImg struct {
	// 图片链接地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 图片放在第几张（多图时可设置）
	Position int64 `json:"position,omitempty" xml:"position,omitempty"`
}

var poolItemImg = sync.Pool{
	New: func() any {
		return new(ItemImg)
	},
}

// GetItemImg() 从对象池中获取ItemImg
func GetItemImg() *ItemImg {
	return poolItemImg.Get().(*ItemImg)
}

// ReleaseItemImg 释放ItemImg
func ReleaseItemImg(v *ItemImg) {
	v.Url = ""
	v.Position = 0
	poolItemImg.Put(v)
}
