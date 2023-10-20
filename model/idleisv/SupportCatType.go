package idleisv

import (
	"sync"
)

// SupportCatType 结构体
type SupportCatType struct {
	// 1 手机，2 潮品，3 家电，8 乐器，9 数码
	SpBizType string `json:"sp_biz_type,omitempty" xml:"sp_biz_type,omitempty"`
}

var poolSupportCatType = sync.Pool{
	New: func() any {
		return new(SupportCatType)
	},
}

// GetSupportCatType() 从对象池中获取SupportCatType
func GetSupportCatType() *SupportCatType {
	return poolSupportCatType.Get().(*SupportCatType)
}

// ReleaseSupportCatType 释放SupportCatType
func ReleaseSupportCatType(v *SupportCatType) {
	v.SpBizType = ""
	poolSupportCatType.Put(v)
}
