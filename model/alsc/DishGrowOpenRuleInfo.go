package alsc

import (
	"sync"
)

// DishGrowOpenRuleInfo 结构体
type DishGrowOpenRuleInfo struct {
	// 规则ID
	DishRuleId string `json:"dish_rule_id,omitempty" xml:"dish_rule_id,omitempty"`
	// 扩展信息
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 会员计划ID
	PlanId string `json:"plan_id,omitempty" xml:"plan_id,omitempty"`
	// 菜品ID
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 升级到的等级id
	ToLevelId string `json:"to_level_id,omitempty" xml:"to_level_id,omitempty"`
	// 外部菜品ID
	OutSkuId string `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	// 是否已删除
	Deleted bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
}

var poolDishGrowOpenRuleInfo = sync.Pool{
	New: func() any {
		return new(DishGrowOpenRuleInfo)
	},
}

// GetDishGrowOpenRuleInfo() 从对象池中获取DishGrowOpenRuleInfo
func GetDishGrowOpenRuleInfo() *DishGrowOpenRuleInfo {
	return poolDishGrowOpenRuleInfo.Get().(*DishGrowOpenRuleInfo)
}

// ReleaseDishGrowOpenRuleInfo 释放DishGrowOpenRuleInfo
func ReleaseDishGrowOpenRuleInfo(v *DishGrowOpenRuleInfo) {
	v.DishRuleId = ""
	v.ExtInfo = ""
	v.GmtCreate = ""
	v.GmtModified = ""
	v.PlanId = ""
	v.SkuId = ""
	v.ToLevelId = ""
	v.OutSkuId = ""
	v.Deleted = false
	poolDishGrowOpenRuleInfo.Put(v)
}
