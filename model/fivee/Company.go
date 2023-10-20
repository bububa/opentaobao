package fivee

import (
	"sync"
)

// Company 结构体
type Company struct {
	// 证照信息
	Licences []Licence `json:"licences,omitempty" xml:"licences>licence,omitempty"`
	// 关系类型：1.制造商 2.供应商
	RelationType []int64 `json:"relation_type,omitempty" xml:"relation_type>int64,omitempty"`
	// 成立时间
	EstablishedDate string `json:"established_date,omitempty" xml:"established_date,omitempty"`
	// 商名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 注册资本
	RegisteredCapital string `json:"registered_capital,omitempty" xml:"registered_capital,omitempty"`
	// 统一社会信用代码
	UniqueCode string `json:"unique_code,omitempty" xml:"unique_code,omitempty"`
}

var poolCompany = sync.Pool{
	New: func() any {
		return new(Company)
	},
}

// GetCompany() 从对象池中获取Company
func GetCompany() *Company {
	return poolCompany.Get().(*Company)
}

// ReleaseCompany 释放Company
func ReleaseCompany(v *Company) {
	v.Licences = v.Licences[:0]
	v.RelationType = v.RelationType[:0]
	v.EstablishedDate = ""
	v.Name = ""
	v.RegisteredCapital = ""
	v.UniqueCode = ""
	poolCompany.Put(v)
}
