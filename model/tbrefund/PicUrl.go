package tbrefund

import (
	"sync"
)

// PicUrl 结构体
type PicUrl struct {
	// 图片链接地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
}

var poolPicUrl = sync.Pool{
	New: func() any {
		return new(PicUrl)
	},
}

// GetPicUrl() 从对象池中获取PicUrl
func GetPicUrl() *PicUrl {
	return poolPicUrl.Get().(*PicUrl)
}

// ReleasePicUrl 释放PicUrl
func ReleasePicUrl(v *PicUrl) {
	v.Url = ""
	poolPicUrl.Put(v)
}
