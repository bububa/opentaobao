package eleenterpriseemployee

import (
	"sync"
)

// CostCenter 结构体
type CostCenter struct {
	// 删除成本中心列表
	DeleteItemIds []string `json:"delete_item_ids,omitempty" xml:"delete_item_ids>string,omitempty"`
	// 新增成本中心列表
	AddItemIds []string `json:"add_item_ids,omitempty" xml:"add_item_ids>string,omitempty"`
}

var poolCostCenter = sync.Pool{
	New: func() any {
		return new(CostCenter)
	},
}

// GetCostCenter() 从对象池中获取CostCenter
func GetCostCenter() *CostCenter {
	return poolCostCenter.Get().(*CostCenter)
}

// ReleaseCostCenter 释放CostCenter
func ReleaseCostCenter(v *CostCenter) {
	v.DeleteItemIds = v.DeleteItemIds[:0]
	v.AddItemIds = v.AddItemIds[:0]
	poolCostCenter.Put(v)
}
