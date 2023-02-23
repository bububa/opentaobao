package alidoc

// SpuDiseaseQueryDto 结构体
type SpuDiseaseQueryDto struct {
	// spu列表，多个逗号
	SpuIds string `json:"spu_ids,omitempty" xml:"spu_ids,omitempty"`
	// 租户
	Tenant string `json:"tenant,omitempty" xml:"tenant,omitempty"`
}
