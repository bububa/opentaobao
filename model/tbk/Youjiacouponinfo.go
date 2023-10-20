package tbk

import (
	"sync"
)

// Youjiacouponinfo 结构体
type Youjiacouponinfo struct {
	// 有价券商品id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品链接
	Url string `json:"url,omitempty" xml:"url,omitempty"`
}

var poolYoujiacouponinfo = sync.Pool{
	New: func() any {
		return new(Youjiacouponinfo)
	},
}

// GetYoujiacouponinfo() 从对象池中获取Youjiacouponinfo
func GetYoujiacouponinfo() *Youjiacouponinfo {
	return poolYoujiacouponinfo.Get().(*Youjiacouponinfo)
}

// ReleaseYoujiacouponinfo 释放Youjiacouponinfo
func ReleaseYoujiacouponinfo(v *Youjiacouponinfo) {
	v.ItemId = ""
	v.Url = ""
	poolYoujiacouponinfo.Put(v)
}
