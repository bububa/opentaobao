package cityretail

// OrderLogisticsResponseDto 
type OrderLogisticsResponseDto struct {

    // 快递员姓名
    
    DelivererName   string `json:"deliverer_name,omitempty" xml:"deliverer_name,omitempty"`
    

    // 快递员电话
    
    DelivererPhone   string `json:"deliverer_phone,omitempty" xml:"deliverer_phone,omitempty"`
    

    // 取件失败原因
    
    FailReason   string `json:"fail_reason,omitempty" xml:"fail_reason,omitempty"`
    

    // 失败码
    
    FailCode   string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
    

    // 状态产生时间
    
    LogisticTime   string `json:"logistic_time,omitempty" xml:"logistic_time,omitempty"`
    

    // 物流状态
    
    LogisticStatus   string `json:"logistic_status,omitempty" xml:"logistic_status,omitempty"`
    

    // 物流状态说明
    
    LogisticStatusName   string `json:"logistic_status_name,omitempty" xml:"logistic_status_name,omitempty"`
    

}
