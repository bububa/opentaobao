package damai

// ThirdStandPushOpenParam 
type ThirdStandPushOpenParam struct {

    // 场次id
    
    PerformId   int64 `json:"perform_id,omitempty" xml:"perform_id,omitempty"`
    

    // 推送时间
    
    PushTime   string `json:"push_time,omitempty" xml:"push_time,omitempty"`
    

    // 看台id
    
    StandId   int64 `json:"stand_id,omitempty" xml:"stand_id,omitempty"`
    

    // 看台名称
    
    StandName   string `json:"stand_name,omitempty" xml:"stand_name,omitempty"`
    

    // 商户密钥
    
    SupplierSecret   string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
    

    // 系统id
    
    SystemId   int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
    

    // 场馆id
    
    VenueId   int64 `json:"venue_id,omitempty" xml:"venue_id,omitempty"`
    

}
