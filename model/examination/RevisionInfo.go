package examination

import (
	"sync"
)

// RevisionInfo 结构体
type RevisionInfo struct {
	// 返回状态码600和700 需要此值
	ReserveDate string `json:"reserve_date,omitempty" xml:"reserve_date,omitempty"`
	// 返回状态码600和700 需要此值
	UniqReserveCode string `json:"uniq_reserve_code,omitempty" xml:"uniq_reserve_code,omitempty"`
}

var poolRevisionInfo = sync.Pool{
	New: func() any {
		return new(RevisionInfo)
	},
}

// GetRevisionInfo() 从对象池中获取RevisionInfo
func GetRevisionInfo() *RevisionInfo {
	return poolRevisionInfo.Get().(*RevisionInfo)
}

// ReleaseRevisionInfo 释放RevisionInfo
func ReleaseRevisionInfo(v *RevisionInfo) {
	v.ReserveDate = ""
	v.UniqReserveCode = ""
	poolRevisionInfo.Put(v)
}
