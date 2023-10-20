package idleisv

import (
	"sync"
)

// HotCompanyList 结构体
type HotCompanyList struct {
	// 快递公司代码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 快递公司名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 快递公司id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolHotCompanyList = sync.Pool{
	New: func() any {
		return new(HotCompanyList)
	},
}

// GetHotCompanyList() 从对象池中获取HotCompanyList
func GetHotCompanyList() *HotCompanyList {
	return poolHotCompanyList.Get().(*HotCompanyList)
}

// ReleaseHotCompanyList 释放HotCompanyList
func ReleaseHotCompanyList(v *HotCompanyList) {
	v.Code = ""
	v.Name = ""
	v.Id = 0
	poolHotCompanyList.Put(v)
}
