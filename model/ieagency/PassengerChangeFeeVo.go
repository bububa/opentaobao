package ieagency

// PassengerChangeFeeVo 
type PassengerChangeFeeVo struct {

    // 改签升舱费(单位:分）
    
    ChangeUpgradeFee   int64 `json:"change_upgrade_fee,omitempty" xml:"change_upgrade_fee,omitempty"`
    

    // 改签服务费(单位:分）
    
    ChangeServiceFee   int64 `json:"change_service_fee,omitempty" xml:"change_service_fee,omitempty"`
    

    // 乘机人ID
    
    PassengerId   int64 `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
    

}
