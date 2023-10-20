package tmallgenie

import (
	"sync"
)

// UnionIdInfo 结构体
type UnionIdInfo struct {
	// 组织id
	OrganizationId string `json:"organization_id,omitempty" xml:"organization_id,omitempty"`
	// 开放unionId
	UnionId string `json:"union_id,omitempty" xml:"union_id,omitempty"`
}

var poolUnionIdInfo = sync.Pool{
	New: func() any {
		return new(UnionIdInfo)
	},
}

// GetUnionIdInfo() 从对象池中获取UnionIdInfo
func GetUnionIdInfo() *UnionIdInfo {
	return poolUnionIdInfo.Get().(*UnionIdInfo)
}

// ReleaseUnionIdInfo 释放UnionIdInfo
func ReleaseUnionIdInfo(v *UnionIdInfo) {
	v.OrganizationId = ""
	v.UnionId = ""
	poolUnionIdInfo.Put(v)
}
