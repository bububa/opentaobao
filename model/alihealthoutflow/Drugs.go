package alihealthoutflow

// Drugs 结构体
type Drugs struct {
	// 规格(非空)
	Spec string `json:"spec,omitempty" xml:"spec,omitempty"`
	// 数量(非空)
	Total string `json:"total,omitempty" xml:"total,omitempty"`
	// 药品名称(可空)
	DrugName string `json:"drug_name,omitempty" xml:"drug_name,omitempty"`
	// 天数(可空)
	Day string `json:"day,omitempty" xml:"day,omitempty"`
	// 频次(非空)
	Frequency string `json:"frequency,omitempty" xml:"frequency,omitempty"`
	// 说明(可空)
	Note string `json:"note,omitempty" xml:"note,omitempty"`
	// 用量(非空)
	Dose string `json:"dose,omitempty" xml:"dose,omitempty"`
	// 药品通用名(非空)
	DrugCommonName string `json:"drug_common_name,omitempty" xml:"drug_common_name,omitempty"`
	// 用量单位(非空)
	DoseUnit string `json:"dose_unit,omitempty" xml:"dose_unit,omitempty"`
	// 总量单位(可空)
	TotalUnit string `json:"total_unit,omitempty" xml:"total_unit,omitempty"`
	// 单价-元(可空)
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 总价-元(可空)
	TotalPrice string `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 用法(非空)
	DoseUsage string `json:"dose_usage,omitempty" xml:"dose_usage,omitempty"`
	// 药品id(drug_id、spuid二选一)
	DrugId string `json:"drug_id,omitempty" xml:"drug_id,omitempty"`
	// 产地id(可空)
	ProduceId string `json:"produce_id,omitempty" xml:"produce_id,omitempty"`
	// spuid (drug_id、spuid二选一) - 纳里必传
	Spuid string `json:"spuid,omitempty" xml:"spuid,omitempty"`
}
