package ascpchannel

import (
	"sync"
)

// Additionalinfo 结构体
type Additionalinfo struct {
	// 附加熟悉字段
	Attribute *Attribute `json:"attribute,omitempty" xml:"attribute,omitempty"`
}

var poolAdditionalinfo = sync.Pool{
	New: func() any {
		return new(Additionalinfo)
	},
}

// GetAdditionalinfo() 从对象池中获取Additionalinfo
func GetAdditionalinfo() *Additionalinfo {
	return poolAdditionalinfo.Get().(*Additionalinfo)
}

// ReleaseAdditionalinfo 释放Additionalinfo
func ReleaseAdditionalinfo(v *Additionalinfo) {
	v.Attribute = nil
	poolAdditionalinfo.Put(v)
}
