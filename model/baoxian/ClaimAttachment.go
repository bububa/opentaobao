package baoxian

import (
	"sync"
)

// ClaimAttachment 结构体
type ClaimAttachment struct {
	// 文件名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 文件路径
	Path string `json:"path,omitempty" xml:"path,omitempty"`
	// 文件描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 文件类型
	FileType string `json:"file_type,omitempty" xml:"file_type,omitempty"`
	// 附件类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 文件大小
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
}

var poolClaimAttachment = sync.Pool{
	New: func() any {
		return new(ClaimAttachment)
	},
}

// GetClaimAttachment() 从对象池中获取ClaimAttachment
func GetClaimAttachment() *ClaimAttachment {
	return poolClaimAttachment.Get().(*ClaimAttachment)
}

// ReleaseClaimAttachment 释放ClaimAttachment
func ReleaseClaimAttachment(v *ClaimAttachment) {
	v.Name = ""
	v.Path = ""
	v.Description = ""
	v.FileType = ""
	v.Type = 0
	v.Size = 0
	poolClaimAttachment.Put(v)
}
