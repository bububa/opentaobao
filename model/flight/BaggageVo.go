package flight

// BaggageVo 
type BaggageVo struct {

    // 是否为全部重量
    
    IsAllWeight   bool `json:"is_all_weight,omitempty" xml:"is_all_weight,omitempty"`
    

    // 行李数量
    
    Pc   int64 `json:"pc,omitempty" xml:"pc,omitempty"`
    

    // 行李重量，单位KG
    
    Weight   int64 `json:"weight,omitempty" xml:"weight,omitempty"`
    

}
