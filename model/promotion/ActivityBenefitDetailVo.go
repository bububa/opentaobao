package promotion

import (
	"sync"
)

// ActivityBenefitDetailVo 结构体
type ActivityBenefitDetailVo struct {
	// 权益类型
	BenefitType string `json:"benefit_type,omitempty" xml:"benefit_type,omitempty"`
	// 权益ID
	BenefitId int64 `json:"benefit_id,omitempty" xml:"benefit_id,omitempty"`
	// 权益标识
	ConfigId int64 `json:"config_id,omitempty" xml:"config_id,omitempty"`
}

var poolActivityBenefitDetailVo = sync.Pool{
	New: func() any {
		return new(ActivityBenefitDetailVo)
	},
}

// GetActivityBenefitDetailVo() 从对象池中获取ActivityBenefitDetailVo
func GetActivityBenefitDetailVo() *ActivityBenefitDetailVo {
	return poolActivityBenefitDetailVo.Get().(*ActivityBenefitDetailVo)
}

// ReleaseActivityBenefitDetailVo 释放ActivityBenefitDetailVo
func ReleaseActivityBenefitDetailVo(v *ActivityBenefitDetailVo) {
	v.BenefitType = ""
	v.BenefitId = 0
	v.ConfigId = 0
	poolActivityBenefitDetailVo.Put(v)
}
