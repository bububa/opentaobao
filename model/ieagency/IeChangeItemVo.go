package ieagency

// IeChangeItemVo 
type IeChangeItemVo struct {

    // 成人改签手续费
    
    AdultServiceFee   int64 `json:"adult_service_fee,omitempty" xml:"adult_service_fee,omitempty"`
    

    // 成人升舱费用
    
    AdultUpgradeFee   int64 `json:"adult_upgrade_fee,omitempty" xml:"adult_upgrade_fee,omitempty"`
    

    // 行李额
    
    BaggageRule   string `json:"baggage_rule,omitempty" xml:"baggage_rule,omitempty"`
    

    // 儿童改签手续费
    
    ChildServiceFee   int64 `json:"child_service_fee,omitempty" xml:"child_service_fee,omitempty"`
    

    // 儿童升舱费用
    
    ChildUpgradeFee   int64 `json:"child_upgrade_fee,omitempty" xml:"child_upgrade_fee,omitempty"`
    

    // 目标行程信息
    
    DestinationItinerarys   []IeChangeItineraryVo `json:"destination_itinerarys,omitempty" xml:"destination_itinerarys,omitempty"`
    

    // 原始改签行程
    
    SourceItinerarys   []IeChangeItineraryVo `json:"source_itinerarys,omitempty" xml:"source_itinerarys,omitempty"`
    

    // 航程类型 0－单程,1－去程,2－回程,3－往返,4－多程
    
    TripType   int64 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
    

    // 婴儿改签升舱费
    
    InfantUpgradeFee   int64 `json:"infant_upgrade_fee,omitempty" xml:"infant_upgrade_fee,omitempty"`
    

    // 婴儿改签费
    
    InfantServiceFee   int64 `json:"infant_service_fee,omitempty" xml:"infant_service_fee,omitempty"`
    

}
