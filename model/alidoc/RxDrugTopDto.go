package alidoc

// RxDrugTopDto 
type RxDrugTopDto struct {

    // 药品用法用量
    
    DrugUsageList   []RxDrugUsageTopDto `json:"drug_usage_list,omitempty" xml:"drug_usage_list,omitempty"`
    

    // 剂型
    
    DoseFrom   string `json:"dose_from,omitempty" xml:"dose_from,omitempty"`
    

    // 数量
    
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
    

    // 规格
    
    Spec   string `json:"spec,omitempty" xml:"spec,omitempty"`
    

    // 药品名称
    
    DrugName   string `json:"drug_name,omitempty" xml:"drug_name,omitempty"`
    

    // 药品Id
    
    DrugId   string `json:"drug_id,omitempty" xml:"drug_id,omitempty"`
    

}
