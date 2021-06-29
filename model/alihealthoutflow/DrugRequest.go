package alihealthoutflow

// DrugRequest 
type DrugRequest struct {

    // 给药疗程单位
    
    TreatmentUnit   string `json:"treatment_unit,omitempty" xml:"treatment_unit,omitempty"`
    

    // 给药疗程值
    
    TreatmentValue   string `json:"treatment_value,omitempty" xml:"treatment_value,omitempty"`
    

    // 给药疗程
    
    TreatmentCourse   string `json:"treatment_course,omitempty" xml:"treatment_course,omitempty"`
    

    // 单剂量单位
    
    SingleDosageUnit   string `json:"single_dosage_unit,omitempty" xml:"single_dosage_unit,omitempty"`
    

    // 单剂量值
    
    SingleDosage   string `json:"single_dosage,omitempty" xml:"single_dosage,omitempty"`
    

    // 包装规格--最小销售单位
    
    PackSpecSaleUnit   string `json:"pack_spec_sale_unit,omitempty" xml:"pack_spec_sale_unit,omitempty"`
    

    // 包装规格--最小使用单位
    
    PackSpecUseUnit   string `json:"pack_spec_use_unit,omitempty" xml:"pack_spec_use_unit,omitempty"`
    

    // 包装规格--包装数量
    
    PackSpecAmount   string `json:"pack_spec_amount,omitempty" xml:"pack_spec_amount,omitempty"`
    

    // 包装单位-装量值
    
    PackUnit   string `json:"pack_unit,omitempty" xml:"pack_unit,omitempty"`
    

    // 包装单位-装量单位
    
    PackQty   string `json:"pack_qty,omitempty" xml:"pack_qty,omitempty"`
    

    // 药品来源（非空）
    
    DrugSource   string `json:"drug_source,omitempty" xml:"drug_source,omitempty"`
    

    // 频次编码
    
    FrequencyCode   string `json:"frequency_code,omitempty" xml:"frequency_code,omitempty"`
    

    // 频次
    
    Frequency   string `json:"frequency,omitempty" xml:"frequency,omitempty"`
    

    // 状态
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    

    // 批准文号
    
    ApprovalNumber   string `json:"approval_number,omitempty" xml:"approval_number,omitempty"`
    

    // 生产厂家（非空）
    
    Manufacture   string `json:"manufacture,omitempty" xml:"manufacture,omitempty"`
    

    // 剂型
    
    DosageForm   string `json:"dosage_form,omitempty" xml:"dosage_form,omitempty"`
    

    // 用法编码
    
    DrugUsageCode   string `json:"drug_usage_code,omitempty" xml:"drug_usage_code,omitempty"`
    

    // 用法
    
    DrugUsage   string `json:"drug_usage,omitempty" xml:"drug_usage,omitempty"`
    

    // his药品编码（非空）
    
    DrugCode   string `json:"drug_code,omitempty" xml:"drug_code,omitempty"`
    

    // 规格（非空）
    
    PackSpec   string `json:"pack_spec,omitempty" xml:"pack_spec,omitempty"`
    

    // 药品通用名（非空）
    
    DrugCommonName   string `json:"drug_common_name,omitempty" xml:"drug_common_name,omitempty"`
    

    // 药品商品名
    
    DrugName   string `json:"drug_name,omitempty" xml:"drug_name,omitempty"`
    

    // 零售价(元)
    
    RetailPrice   string `json:"retail_price,omitempty" xml:"retail_price,omitempty"`
    

    // 药品本位码
    
    DrugBasicCode   string `json:"drug_basic_code,omitempty" xml:"drug_basic_code,omitempty"`
    

    // 医保编号
    
    MedicareNo   string `json:"medicare_no,omitempty" xml:"medicare_no,omitempty"`
    

    // 频次数量
    
    FreqTimes   string `json:"freq_times,omitempty" xml:"freq_times,omitempty"`
    

}
