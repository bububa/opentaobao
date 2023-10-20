package btrip

import (
	"sync"
)

// OpenCostCenterDo 结构体
type OpenCostCenterDo struct {
	// 成本中心编号
	Number string `json:"number,omitempty" xml:"number,omitempty"`
	// 成本中心名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 企业id
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 商旅企业id
	Corpid string `json:"corpid,omitempty" xml:"corpid,omitempty"`
	// 商旅成本中心id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolOpenCostCenterDo = sync.Pool{
	New: func() any {
		return new(OpenCostCenterDo)
	},
}

// GetOpenCostCenterDo() 从对象池中获取OpenCostCenterDo
func GetOpenCostCenterDo() *OpenCostCenterDo {
	return poolOpenCostCenterDo.Get().(*OpenCostCenterDo)
}

// ReleaseOpenCostCenterDo 释放OpenCostCenterDo
func ReleaseOpenCostCenterDo(v *OpenCostCenterDo) {
	v.Number = ""
	v.Name = ""
	v.CorpId = ""
	v.Corpid = ""
	v.Id = 0
	poolOpenCostCenterDo.Put(v)
}
