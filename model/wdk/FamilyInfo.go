package wdk

import (
	"sync"
)

// FamilyInfo 结构体
type FamilyInfo struct {
	// 手机号码
	Contact string `json:"contact,omitempty" xml:"contact,omitempty"`
	// 家庭成员姓名
	FamilyName string `json:"family_name,omitempty" xml:"family_name,omitempty"`
	// 职位
	Post string `json:"post,omitempty" xml:"post,omitempty"`
	// 关系
	Relationship string `json:"relationship,omitempty" xml:"relationship,omitempty"`
	// 工作单位
	WorkCompany string `json:"work_company,omitempty" xml:"work_company,omitempty"`
}

var poolFamilyInfo = sync.Pool{
	New: func() any {
		return new(FamilyInfo)
	},
}

// GetFamilyInfo() 从对象池中获取FamilyInfo
func GetFamilyInfo() *FamilyInfo {
	return poolFamilyInfo.Get().(*FamilyInfo)
}

// ReleaseFamilyInfo 释放FamilyInfo
func ReleaseFamilyInfo(v *FamilyInfo) {
	v.Contact = ""
	v.FamilyName = ""
	v.Post = ""
	v.Relationship = ""
	v.WorkCompany = ""
	poolFamilyInfo.Put(v)
}
