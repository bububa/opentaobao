package miniappopen

import (
	"sync"
)

// DistributionOrderBindTargetEntityOpenResultDto 结构体
type DistributionOrderBindTargetEntityOpenResultDto struct {
	// 商品列表的绑定结果
	BindResultList []DistributionOrderBindBaseDto `json:"bind_result_list,omitempty" xml:"bind_result_list>distribution_order_bind_base_dto,omitempty"`
	// 绑定的目标url
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 绑定投放的场景
	SceneName string `json:"scene_name,omitempty" xml:"scene_name,omitempty"`
	// 绑定的投放计划id
	DistributeId int64 `json:"distribute_id,omitempty" xml:"distribute_id,omitempty"`
}

var poolDistributionOrderBindTargetEntityOpenResultDto = sync.Pool{
	New: func() any {
		return new(DistributionOrderBindTargetEntityOpenResultDto)
	},
}

// GetDistributionOrderBindTargetEntityOpenResultDto() 从对象池中获取DistributionOrderBindTargetEntityOpenResultDto
func GetDistributionOrderBindTargetEntityOpenResultDto() *DistributionOrderBindTargetEntityOpenResultDto {
	return poolDistributionOrderBindTargetEntityOpenResultDto.Get().(*DistributionOrderBindTargetEntityOpenResultDto)
}

// ReleaseDistributionOrderBindTargetEntityOpenResultDto 释放DistributionOrderBindTargetEntityOpenResultDto
func ReleaseDistributionOrderBindTargetEntityOpenResultDto(v *DistributionOrderBindTargetEntityOpenResultDto) {
	v.BindResultList = v.BindResultList[:0]
	v.Url = ""
	v.SceneName = ""
	v.DistributeId = 0
	poolDistributionOrderBindTargetEntityOpenResultDto.Put(v)
}
