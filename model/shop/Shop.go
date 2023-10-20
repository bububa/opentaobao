package shop

import (
	"sync"
)

// Shop 结构体
type Shop struct {
	// 店铺标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 店标地址。返回相对路径，可以用&#34;http://logo.taobao.com/shop-logo&#34;来拼接成绝对路径
	PicPath string `json:"pic_path,omitempty" xml:"pic_path,omitempty"`
	// 最后修改时间。格式：yyyy-MM-dd HH:mm:ss
	Modified string `json:"modified,omitempty" xml:"modified,omitempty"`
	// 店铺编号
	Sid int64 `json:"sid,omitempty" xml:"sid,omitempty"`
}

var poolShop = sync.Pool{
	New: func() any {
		return new(Shop)
	},
}

// GetShop() 从对象池中获取Shop
func GetShop() *Shop {
	return poolShop.Get().(*Shop)
}

// ReleaseShop 释放Shop
func ReleaseShop(v *Shop) {
	v.Title = ""
	v.PicPath = ""
	v.Modified = ""
	v.Sid = 0
	poolShop.Put(v)
}
