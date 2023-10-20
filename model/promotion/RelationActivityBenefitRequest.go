package promotion

import (
	"sync"
)

// RelationActivityBenefitRequest 结构体
type RelationActivityBenefitRequest struct {
	// 活动关联的权益信息，可以从权益选择器API中获取
	AddDetailVos []ActivityBenefitDetailVo `json:"add_detail_vos,omitempty" xml:"add_detail_vos>activity_benefit_detail_vo,omitempty"`
	// 同步权益活动的概述信息，用于卖家后台查看
	BenefitActivityVo *BenefitActivityVo `json:"benefit_activity_vo,omitempty" xml:"benefit_activity_vo,omitempty"`
}

var poolRelationActivityBenefitRequest = sync.Pool{
	New: func() any {
		return new(RelationActivityBenefitRequest)
	},
}

// GetRelationActivityBenefitRequest() 从对象池中获取RelationActivityBenefitRequest
func GetRelationActivityBenefitRequest() *RelationActivityBenefitRequest {
	return poolRelationActivityBenefitRequest.Get().(*RelationActivityBenefitRequest)
}

// ReleaseRelationActivityBenefitRequest 释放RelationActivityBenefitRequest
func ReleaseRelationActivityBenefitRequest(v *RelationActivityBenefitRequest) {
	v.AddDetailVos = v.AddDetailVos[:0]
	v.BenefitActivityVo = nil
	poolRelationActivityBenefitRequest.Put(v)
}
