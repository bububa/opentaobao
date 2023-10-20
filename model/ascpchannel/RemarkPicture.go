package ascpchannel

import (
	"sync"
)

// RemarkPicture 结构体
type RemarkPicture struct {
	// 备注图片英文逗号分隔
	DownloadUrl string `json:"download_url,omitempty" xml:"download_url,omitempty"`
}

var poolRemarkPicture = sync.Pool{
	New: func() any {
		return new(RemarkPicture)
	},
}

// GetRemarkPicture() 从对象池中获取RemarkPicture
func GetRemarkPicture() *RemarkPicture {
	return poolRemarkPicture.Get().(*RemarkPicture)
}

// ReleaseRemarkPicture 释放RemarkPicture
func ReleaseRemarkPicture(v *RemarkPicture) {
	v.DownloadUrl = ""
	poolRemarkPicture.Put(v)
}
