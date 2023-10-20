package promotion

import (
	"sync"
)

// BenefitActivityVo 结构体
type BenefitActivityVo struct {
	// ISV活动的具体地址
	ActivityUrl string `json:"activity_url,omitempty" xml:"activity_url,omitempty"`
	// 活动描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 活动结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 活动名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 活动的开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 活动类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
}

var poolBenefitActivityVo = sync.Pool{
	New: func() any {
		return new(BenefitActivityVo)
	},
}

// GetBenefitActivityVo() 从对象池中获取BenefitActivityVo
func GetBenefitActivityVo() *BenefitActivityVo {
	return poolBenefitActivityVo.Get().(*BenefitActivityVo)
}

// ReleaseBenefitActivityVo 释放BenefitActivityVo
func ReleaseBenefitActivityVo(v *BenefitActivityVo) {
	v.ActivityUrl = ""
	v.Desc = ""
	v.EndTime = ""
	v.Name = ""
	v.StartTime = ""
	v.Type = ""
	poolBenefitActivityVo.Put(v)
}
