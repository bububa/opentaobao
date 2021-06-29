package alihealthoutflow

// DrugDTO 
type DrugDTO struct {
    // 用药单位
    DoseUnit   string `json:"dose_unit,omitempty" xml:"dose_unit,omitempty"`
    // 药品名称
    DrugName   string `json:"drug_name,omitempty" xml:"drug_name,omitempty"`
    // 总量
    Total   string `json:"total,omitempty" xml:"total,omitempty"`
    // 用药数量
    Dose   string `json:"dose,omitempty" xml:"dose,omitempty"`
    // his频次编码
    Frequency   string `json:"frequency,omitempty" xml:"frequency,omitempty"`
    // 用法用量+医嘱
    DoseUsageAdvice   string `json:"dose_usage_advice,omitempty" xml:"dose_usage_advice,omitempty"`
    // 药品厂家
    Manufactures   string `json:"manufactures,omitempty" xml:"manufactures,omitempty"`
    // 规格
    Spec   string `json:"spec,omitempty" xml:"spec,omitempty"`
    // 商品名称
    ProdName   string `json:"prod_name,omitempty" xml:"prod_name,omitempty"`
    // 药品编码
    DrugCode   string `json:"drug_code,omitempty" xml:"drug_code,omitempty"`
    // 药品数量（必选）
    Count   string `json:"count,omitempty" xml:"count,omitempty"`
    // 核销药品单价
    Price   string `json:"price,omitempty" xml:"price,omitempty"`
}
