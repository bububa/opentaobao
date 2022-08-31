package alihealth2

// Drug 结构体
type Drug struct {
	// 成分
	Ingredients []Ingredients `json:"ingredients,omitempty" xml:"ingredients>ingredients,omitempty"`
	// 通用名
	CommonName string `json:"common_name,omitempty" xml:"common_name,omitempty"`
	// 租户的药品唯一标识别
	TenantSpuId string `json:"tenant_spu_id,omitempty" xml:"tenant_spu_id,omitempty"`
	// 生产厂家
	Producer string `json:"producer,omitempty" xml:"producer,omitempty"`
	// 健康标品库的SpuId
	SpuId string `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	// 最小使用单位数值剂量单位（mg）
	UnitOfPerUnit string `json:"unit_of_per_unit,omitempty" xml:"unit_of_per_unit,omitempty"`
	// 国药准字号
	ApprovalNumber string `json:"approval_number,omitempty" xml:"approval_number,omitempty"`
	// 药品类型（WESTERN_PRESCRIPTION_MEDICINE(1,&#34;处方药-西药&#34;),     CHINESE_PRESCRIPTION_MEDICINE(2, &#34;处方药-中成药&#34;),     WESTERN_OTC(3, &#34;OTC-西药&#34;),     CHINESE_OTC(4, &#34;OTC-中成药&#34;),     PIECES(5, &#34;饮片&#34;)）
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 规格详情
	Norm string `json:"norm,omitempty" xml:"norm,omitempty"`
	// 装量值(例：24粒)
	QuantityPerPack int64 `json:"quantity_per_pack,omitempty" xml:"quantity_per_pack,omitempty"`
	// 最小使用单位数值（例：10）
	QuantityPerUnit int64 `json:"quantity_per_unit,omitempty" xml:"quantity_per_unit,omitempty"`
}
