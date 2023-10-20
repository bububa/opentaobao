package wdk

import (
	"sync"
)

// RelatedPartyInfo 结构体
type RelatedPartyInfo struct {
	// 所在部门
	Department string `json:"department,omitempty" xml:"department,omitempty"`
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 职位
	Post string `json:"post,omitempty" xml:"post,omitempty"`
	// 关联方类型
	RelatedPartyType string `json:"related_party_type,omitempty" xml:"related_party_type,omitempty"`
	// 关系
	Relation string `json:"relation,omitempty" xml:"relation,omitempty"`
}

var poolRelatedPartyInfo = sync.Pool{
	New: func() any {
		return new(RelatedPartyInfo)
	},
}

// GetRelatedPartyInfo() 从对象池中获取RelatedPartyInfo
func GetRelatedPartyInfo() *RelatedPartyInfo {
	return poolRelatedPartyInfo.Get().(*RelatedPartyInfo)
}

// ReleaseRelatedPartyInfo 释放RelatedPartyInfo
func ReleaseRelatedPartyInfo(v *RelatedPartyInfo) {
	v.Department = ""
	v.Name = ""
	v.Post = ""
	v.RelatedPartyType = ""
	v.Relation = ""
	poolRelatedPartyInfo.Put(v)
}
