package baoxian

import (
	"sync"
)

// InsAttachmentUploadVo 结构体
type InsAttachmentUploadVo struct {
	// ossPath
	OssPath string `json:"oss_path,omitempty" xml:"oss_path,omitempty"`
	// eTag
	ETag string `json:"e_tag,omitempty" xml:"e_tag,omitempty"`
	// size
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
}

var poolInsAttachmentUploadVo = sync.Pool{
	New: func() any {
		return new(InsAttachmentUploadVo)
	},
}

// GetInsAttachmentUploadVo() 从对象池中获取InsAttachmentUploadVo
func GetInsAttachmentUploadVo() *InsAttachmentUploadVo {
	return poolInsAttachmentUploadVo.Get().(*InsAttachmentUploadVo)
}

// ReleaseInsAttachmentUploadVo 释放InsAttachmentUploadVo
func ReleaseInsAttachmentUploadVo(v *InsAttachmentUploadVo) {
	v.OssPath = ""
	v.ETag = ""
	v.Size = 0
	poolInsAttachmentUploadVo.Put(v)
}
