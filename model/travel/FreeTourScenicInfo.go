package travel

// FreeTourScenicInfo 
/* model for simplify = false
type FreeTourScenicInfo struct {

    // 必填，景点名称
    
    CnName   string `json:"cn_name,omitempty"`
    

    // 必填，门票类型
    
    TicketType   string `json:"ticket_type,omitempty"`
    

    // 必填，景点所在城市
    
    City   string `json:"city,omitempty"`
    

}
*/

// FreeTourScenicInfo 
type FreeTourScenicInfo struct {

    // 必填，景点名称
    CnName   string `json:"cn_name,omitempty"`

    // 必填，门票类型
    TicketType   string `json:"ticket_type,omitempty"`

    // 必填，景点所在城市
    City   string `json:"city,omitempty"`

}
