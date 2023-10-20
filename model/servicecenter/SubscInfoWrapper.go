package servicecenter

import (
	"sync"
)

// SubscInfoWrapper 结构体
type SubscInfoWrapper struct {
	// 需求订购信息
	SubscInfoList []SubscInfo `json:"subsc_info_list,omitempty" xml:"subsc_info_list>subsc_info,omitempty"`
	// 总量
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

var poolSubscInfoWrapper = sync.Pool{
	New: func() any {
		return new(SubscInfoWrapper)
	},
}

// GetSubscInfoWrapper() 从对象池中获取SubscInfoWrapper
func GetSubscInfoWrapper() *SubscInfoWrapper {
	return poolSubscInfoWrapper.Get().(*SubscInfoWrapper)
}

// ReleaseSubscInfoWrapper 释放SubscInfoWrapper
func ReleaseSubscInfoWrapper(v *SubscInfoWrapper) {
	v.SubscInfoList = v.SubscInfoList[:0]
	v.TotalCount = 0
	poolSubscInfoWrapper.Put(v)
}
