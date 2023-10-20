package idleisv

import (
	"sync"
)

// CompanyList 结构体
type CompanyList struct {
	// 快递公司代码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 快递公司名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 快递公司id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolCompanyList = sync.Pool{
	New: func() any {
		return new(CompanyList)
	},
}

// GetCompanyList() 从对象池中获取CompanyList
func GetCompanyList() *CompanyList {
	return poolCompanyList.Get().(*CompanyList)
}

// ReleaseCompanyList 释放CompanyList
func ReleaseCompanyList(v *CompanyList) {
	v.Code = ""
	v.Name = ""
	v.Id = 0
	poolCompanyList.Put(v)
}
