package alidoc

// DrugDiseaseMappingDto 结构体
type DrugDiseaseMappingDto struct {
	// 诊断信息
	DiseaseList []DiseaseInfo `json:"disease_list,omitempty" xml:"disease_list>disease_info,omitempty"`
	// spuid
	SpuId string `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
}
