package product

import (
	"sync"
)

// UploadImageResp 结构体
type UploadImageResp struct {
	// 返回的图片地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 图片哈希值
	Hash string `json:"hash,omitempty" xml:"hash,omitempty"`
	// 图片高度
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
	// 图片宽度
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
}

var poolUploadImageResp = sync.Pool{
	New: func() any {
		return new(UploadImageResp)
	},
}

// GetUploadImageResp() 从对象池中获取UploadImageResp
func GetUploadImageResp() *UploadImageResp {
	return poolUploadImageResp.Get().(*UploadImageResp)
}

// ReleaseUploadImageResp 释放UploadImageResp
func ReleaseUploadImageResp(v *UploadImageResp) {
	v.Url = ""
	v.Hash = ""
	v.Height = 0
	v.Width = 0
	poolUploadImageResp.Put(v)
}
