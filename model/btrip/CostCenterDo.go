package btrip

import (
	"sync"
)

// CostCenterDo 结构体
type CostCenterDo struct {
	// 成本中心名称
	CostCenterTitle string `json:"cost_center_title,omitempty" xml:"cost_center_title,omitempty"`
	// 成本中心编码
	CostCenterNumber string `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	// 第三方成本中心id
	ThirdCostCenterId string `json:"third_cost_center_id,omitempty" xml:"third_cost_center_id,omitempty"`
	// 成本中心ID
	CostCenterId int64 `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
}

var poolCostCenterDo = sync.Pool{
	New: func() any {
		return new(CostCenterDo)
	},
}

// GetCostCenterDo() 从对象池中获取CostCenterDo
func GetCostCenterDo() *CostCenterDo {
	return poolCostCenterDo.Get().(*CostCenterDo)
}

// ReleaseCostCenterDo 释放CostCenterDo
func ReleaseCostCenterDo(v *CostCenterDo) {
	v.CostCenterTitle = ""
	v.CostCenterNumber = ""
	v.ThirdCostCenterId = ""
	v.CostCenterId = 0
	poolCostCenterDo.Put(v)
}
