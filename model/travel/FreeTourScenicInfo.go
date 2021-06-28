package travel

// FreeTourScenicInfo 
type FreeTourScenicInfo struct {

    // 必填，景点名称
    
    CnName   string `json:"cn_name,omitempty" xml:"cn_name,omitempty"`
    

    // 必填，门票类型
    
    TicketType   string `json:"ticket_type,omitempty" xml:"ticket_type,omitempty"`
    

    // 必填，景点所在城市
    
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    

}
