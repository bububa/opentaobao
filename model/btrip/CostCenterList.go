package btrip

import (
	"sync"
)

// CostCenterList 结构体
type CostCenterList struct {
	// corpId
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// number
	Number string `json:"number,omitempty" xml:"number,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolCostCenterList = sync.Pool{
	New: func() any {
		return new(CostCenterList)
	},
}

// GetCostCenterList() 从对象池中获取CostCenterList
func GetCostCenterList() *CostCenterList {
	return poolCostCenterList.Get().(*CostCenterList)
}

// ReleaseCostCenterList 释放CostCenterList
func ReleaseCostCenterList(v *CostCenterList) {
	v.CorpId = ""
	v.Number = ""
	v.Name = ""
	v.Id = 0
	poolCostCenterList.Put(v)
}
