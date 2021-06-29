package alihealth2

// Drug 
type Drug struct {
    // 通用名
    CommonName   string `json:"common_name,omitempty" xml:"common_name,omitempty"`
    // 装量值(例：24粒)
    QuantityPerPack   int64 `json:"quantity_per_pack,omitempty" xml:"quantity_per_pack,omitempty"`
    // 租户的药品唯一标识别
    TenantSpuId   string `json:"tenant_spu_id,omitempty" xml:"tenant_spu_id,omitempty"`
    // 生产厂家
    Producer   string `json:"producer,omitempty" xml:"producer,omitempty"`
    // 成分
    Ingredients   []Ingredients `json:"ingredients,omitempty" xml:"ingredients>ingredients,omitempty"`
    // 健康标品库的SpuId
    SpuId   string `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
    // 最小使用单位数值剂量单位（mg）
    UnitOfPerUnit   string `json:"unit_of_per_unit,omitempty" xml:"unit_of_per_unit,omitempty"`
    // 最小使用单位数值（例：10）
    QuantityPerUnit   int64 `json:"quantity_per_unit,omitempty" xml:"quantity_per_unit,omitempty"`
    // 国药准字号
    ApprovalNumber   string `json:"approval_number,omitempty" xml:"approval_number,omitempty"`
    // 药品类型（WESTERN_PRESCRIPTION_MEDICINE(1,"处方药-西药"),     CHINESE_PRESCRIPTION_MEDICINE(2, "处方药-中成药"),     WESTERN_OTC(3, "OTC-西药"),     CHINESE_OTC(4, "OTC-中成药"),     PIECES(5, "饮片")）
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
    // 规格详情
    Norm   string `json:"norm,omitempty" xml:"norm,omitempty"`
}
