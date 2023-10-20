package damai

import (
	"sync"
)

// Extrainfomap 结构体
type Extrainfomap struct {
	// 购票链接
	BuyUrl string `json:"buy_url,omitempty" xml:"buy_url,omitempty"`
}

var poolExtrainfomap = sync.Pool{
	New: func() any {
		return new(Extrainfomap)
	},
}

// GetExtrainfomap() 从对象池中获取Extrainfomap
func GetExtrainfomap() *Extrainfomap {
	return poolExtrainfomap.Get().(*Extrainfomap)
}

// ReleaseExtrainfomap 释放Extrainfomap
func ReleaseExtrainfomap(v *Extrainfomap) {
	v.BuyUrl = ""
	poolExtrainfomap.Put(v)
}
