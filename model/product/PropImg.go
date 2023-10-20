package product

import (
	"sync"
)

// PropImg 结构体
type PropImg struct {
	// 图片链接地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 图片所对应的属性组合的字符串
	Properties string `json:"properties,omitempty" xml:"properties,omitempty"`
	// 图片放在第几张（多图时可设置）
	Position int64 `json:"position,omitempty" xml:"position,omitempty"`
}

var poolPropImg = sync.Pool{
	New: func() any {
		return new(PropImg)
	},
}

// GetPropImg() 从对象池中获取PropImg
func GetPropImg() *PropImg {
	return poolPropImg.Get().(*PropImg)
}

// ReleasePropImg 释放PropImg
func ReleasePropImg(v *PropImg) {
	v.Url = ""
	v.Properties = ""
	v.Position = 0
	poolPropImg.Put(v)
}
