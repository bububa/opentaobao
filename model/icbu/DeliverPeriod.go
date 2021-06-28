package icbu

// DeliverPeriod 
type DeliverPeriod struct {

    // 预计需要发货时间
    
    ProcessPeriod   int64 `json:"process_period,omitempty" xml:"process_period,omitempty"`
    

    // 数量
    
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    

}
