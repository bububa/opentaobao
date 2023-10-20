package promotion

import (
	"sync"
)

// DeleteActivityBenefitDetailVo 结构体
type DeleteActivityBenefitDetailVo struct {
	// 活动关联权益后生产的详情ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolDeleteActivityBenefitDetailVo = sync.Pool{
	New: func() any {
		return new(DeleteActivityBenefitDetailVo)
	},
}

// GetDeleteActivityBenefitDetailVo() 从对象池中获取DeleteActivityBenefitDetailVo
func GetDeleteActivityBenefitDetailVo() *DeleteActivityBenefitDetailVo {
	return poolDeleteActivityBenefitDetailVo.Get().(*DeleteActivityBenefitDetailVo)
}

// ReleaseDeleteActivityBenefitDetailVo 释放DeleteActivityBenefitDetailVo
func ReleaseDeleteActivityBenefitDetailVo(v *DeleteActivityBenefitDetailVo) {
	v.Id = 0
	poolDeleteActivityBenefitDetailVo.Put(v)
}
