package alidoc

import (
	"sync"
)

// SpuDiseaseQueryDto 结构体
type SpuDiseaseQueryDto struct {
	// spu列表，多个逗号
	SpuIds string `json:"spu_ids,omitempty" xml:"spu_ids,omitempty"`
	// 租户
	Tenant string `json:"tenant,omitempty" xml:"tenant,omitempty"`
}

var poolSpuDiseaseQueryDto = sync.Pool{
	New: func() any {
		return new(SpuDiseaseQueryDto)
	},
}

// GetSpuDiseaseQueryDto() 从对象池中获取SpuDiseaseQueryDto
func GetSpuDiseaseQueryDto() *SpuDiseaseQueryDto {
	return poolSpuDiseaseQueryDto.Get().(*SpuDiseaseQueryDto)
}

// ReleaseSpuDiseaseQueryDto 释放SpuDiseaseQueryDto
func ReleaseSpuDiseaseQueryDto(v *SpuDiseaseQueryDto) {
	v.SpuIds = ""
	v.Tenant = ""
	poolSpuDiseaseQueryDto.Put(v)
}
