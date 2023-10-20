package crm

import (
	"sync"
)

// GradePromotion 结构体
type GradePromotion struct {
	// 店铺客户、普通会员 、高级会员、VIP会员、至尊VIP
	CurGradeName string `json:"cur_grade_name,omitempty" xml:"cur_grade_name,omitempty"`
	// 买家会员级别  0:店铺客户  1：普通会员 2：高级会员 3：VIP会员 4：至尊VIP
	CurGrade string `json:"cur_grade,omitempty" xml:"cur_grade,omitempty"`
	// 该等级对应的下一等级1:普通会员  2：高级会员 3：VIP会员 4：至尊VIP 0：后续等级都未启用或没有下一等级
	NextGradeName string `json:"next_grade_name,omitempty" xml:"next_grade_name,omitempty"`
	// 普通会员、高级会员、VIP会员、至尊VIP。空的时候代表后续等级未启用或没有下一等级
	NextGrade string `json:"next_grade,omitempty" xml:"next_grade,omitempty"`
	// 升级到下一个级别的需要的交易量
	NextUpgradeCount int64 `json:"next_upgrade_count,omitempty" xml:"next_upgrade_count,omitempty"`
	// 升级到下一个级别的需要的交易额，单位：分
	NextUpgradeAmount int64 `json:"next_upgrade_amount,omitempty" xml:"next_upgrade_amount,omitempty"`
	// 会员级别折扣率没有小数，990代表9.9折
	Discount int64 `json:"discount,omitempty" xml:"discount,omitempty"`
}

var poolGradePromotion = sync.Pool{
	New: func() any {
		return new(GradePromotion)
	},
}

// GetGradePromotion() 从对象池中获取GradePromotion
func GetGradePromotion() *GradePromotion {
	return poolGradePromotion.Get().(*GradePromotion)
}

// ReleaseGradePromotion 释放GradePromotion
func ReleaseGradePromotion(v *GradePromotion) {
	v.CurGradeName = ""
	v.CurGrade = ""
	v.NextGradeName = ""
	v.NextGrade = ""
	v.NextUpgradeCount = 0
	v.NextUpgradeAmount = 0
	v.Discount = 0
	poolGradePromotion.Put(v)
}
