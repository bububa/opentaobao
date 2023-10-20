package promotion

import (
	"sync"
)

// UpdateBenefitActivityVo 结构体
type UpdateBenefitActivityVo struct {
	// ISV活动的活动地址
	ActivityUrl string `json:"activity_url,omitempty" xml:"activity_url,omitempty"`
	// 活动描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 活动名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// ISV活动关联权益后获得的关联ID
	RelationId int64 `json:"relation_id,omitempty" xml:"relation_id,omitempty"`
}

var poolUpdateBenefitActivityVo = sync.Pool{
	New: func() any {
		return new(UpdateBenefitActivityVo)
	},
}

// GetUpdateBenefitActivityVo() 从对象池中获取UpdateBenefitActivityVo
func GetUpdateBenefitActivityVo() *UpdateBenefitActivityVo {
	return poolUpdateBenefitActivityVo.Get().(*UpdateBenefitActivityVo)
}

// ReleaseUpdateBenefitActivityVo 释放UpdateBenefitActivityVo
func ReleaseUpdateBenefitActivityVo(v *UpdateBenefitActivityVo) {
	v.ActivityUrl = ""
	v.Desc = ""
	v.Name = ""
	v.RelationId = 0
	poolUpdateBenefitActivityVo.Put(v)
}
