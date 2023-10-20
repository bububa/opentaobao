package mos

import (
	"sync"
)

// OperatorUserInfo 结构体
type OperatorUserInfo struct {
	// 阿里工号
	AliWorkNo string `json:"ali_work_no,omitempty" xml:"ali_work_no,omitempty"`
	// 银泰工号
	CompWorkNo string `json:"comp_work_no,omitempty" xml:"comp_work_no,omitempty"`
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// operator id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 淘宝id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

var poolOperatorUserInfo = sync.Pool{
	New: func() any {
		return new(OperatorUserInfo)
	},
}

// GetOperatorUserInfo() 从对象池中获取OperatorUserInfo
func GetOperatorUserInfo() *OperatorUserInfo {
	return poolOperatorUserInfo.Get().(*OperatorUserInfo)
}

// ReleaseOperatorUserInfo 释放OperatorUserInfo
func ReleaseOperatorUserInfo(v *OperatorUserInfo) {
	v.AliWorkNo = ""
	v.CompWorkNo = ""
	v.Name = ""
	v.Id = 0
	v.UserId = 0
	poolOperatorUserInfo.Put(v)
}
