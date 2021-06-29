package alihealthoutflow

// DrugUsageVo 
type DrugUsageVo struct {

    // 每次用量单位
    
    MeasureUnit   string `json:"measure_unit,omitempty" xml:"measure_unit,omitempty"`
    

    // 用法
    
    DrugUsage   string `json:"drug_usage,omitempty" xml:"drug_usage,omitempty"`
    

    // 天数
    
    Days   string `json:"days,omitempty" xml:"days,omitempty"`
    

    // 频次值单位
    
    FrequencyUnit   string `json:"frequency_unit,omitempty" xml:"frequency_unit,omitempty"`
    

    // 频次值
    
    FrequencyValue   string `json:"frequency_value,omitempty" xml:"frequency_value,omitempty"`
    

    // 频次
    
    Frequency   string `json:"frequency,omitempty" xml:"frequency,omitempty"`
    

    // 每次你用量
    
    DoseValue   string `json:"dose_value,omitempty" xml:"dose_value,omitempty"`
    

}
