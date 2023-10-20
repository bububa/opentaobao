package fenxiao

import (
	"sync"
)

// TopMemoAttachment 结构体
type TopMemoAttachment struct {
	// 附件地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 附件名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolTopMemoAttachment = sync.Pool{
	New: func() any {
		return new(TopMemoAttachment)
	},
}

// GetTopMemoAttachment() 从对象池中获取TopMemoAttachment
func GetTopMemoAttachment() *TopMemoAttachment {
	return poolTopMemoAttachment.Get().(*TopMemoAttachment)
}

// ReleaseTopMemoAttachment 释放TopMemoAttachment
func ReleaseTopMemoAttachment(v *TopMemoAttachment) {
	v.Url = ""
	v.Name = ""
	poolTopMemoAttachment.Put(v)
}
