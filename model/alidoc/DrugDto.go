package alidoc

// DrugDto 结构体
type DrugDto struct {
	// 药品用法用量
	DrugUsageList []DrugUsageDto `json:"drug_usage_list,omitempty" xml:"drug_usage_list>drug_usage_dto,omitempty"`
	// 药品Id
	DrugId string `json:"drug_id,omitempty" xml:"drug_id,omitempty"`
	// 药品名称
	DrugName string `json:"drug_name,omitempty" xml:"drug_name,omitempty"`
	// 规格
	Spec string `json:"spec,omitempty" xml:"spec,omitempty"`
	// 数量
	Total string `json:"total,omitempty" xml:"total,omitempty"`
	// 剂型
	DoseFrom string `json:"dose_from,omitempty" xml:"dose_from,omitempty"`
}
