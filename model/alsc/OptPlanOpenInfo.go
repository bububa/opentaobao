package alsc

import (
	"sync"
)

// OptPlanOpenInfo 结构体
type OptPlanOpenInfo struct {
	// 运营计划id
	OptPlanId string `json:"opt_plan_id,omitempty" xml:"opt_plan_id,omitempty"`
	// 运营计划名称
	OptPlanName string `json:"opt_plan_name,omitempty" xml:"opt_plan_name,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 储值规则Id
	RechargeRuleId string `json:"recharge_rule_id,omitempty" xml:"recharge_rule_id,omitempty"`
	// 储值规则名称
	RechargeRuleName string `json:"recharge_rule_name,omitempty" xml:"recharge_rule_name,omitempty"`
	// 积分规则Id
	PointRuleId string `json:"point_rule_id,omitempty" xml:"point_rule_id,omitempty"`
	// 积分规则名称
	PointRuleName string `json:"point_rule_name,omitempty" xml:"point_rule_name,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 更新时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 创建者
	CreateBy string `json:"create_by,omitempty" xml:"create_by,omitempty"`
	// 更新者
	UpdateBy string `json:"update_by,omitempty" xml:"update_by,omitempty"`
	// 运营计划类型
	OptPlanType int64 `json:"opt_plan_type,omitempty" xml:"opt_plan_type,omitempty"`
	// 适用门店数
	ApplyShopCount int64 `json:"apply_shop_count,omitempty" xml:"apply_shop_count,omitempty"`
	// 门店组信息
	ShopGroupInfo *ShopGroupOpenInfo `json:"shop_group_info,omitempty" xml:"shop_group_info,omitempty"`
	// 逻辑删除标志
	Deleted bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
}

var poolOptPlanOpenInfo = sync.Pool{
	New: func() any {
		return new(OptPlanOpenInfo)
	},
}

// GetOptPlanOpenInfo() 从对象池中获取OptPlanOpenInfo
func GetOptPlanOpenInfo() *OptPlanOpenInfo {
	return poolOptPlanOpenInfo.Get().(*OptPlanOpenInfo)
}

// ReleaseOptPlanOpenInfo 释放OptPlanOpenInfo
func ReleaseOptPlanOpenInfo(v *OptPlanOpenInfo) {
	v.OptPlanId = ""
	v.OptPlanName = ""
	v.Remark = ""
	v.RechargeRuleId = ""
	v.RechargeRuleName = ""
	v.PointRuleId = ""
	v.PointRuleName = ""
	v.GmtCreate = ""
	v.GmtModified = ""
	v.CreateBy = ""
	v.UpdateBy = ""
	v.OptPlanType = 0
	v.ApplyShopCount = 0
	v.ShopGroupInfo = nil
	v.Deleted = false
	poolOptPlanOpenInfo.Put(v)
}
