package refund

// RefundRemindTimeout 
/* model for simplify = false
type RefundRemindTimeout struct {

    // 提醒的类型（退款详情中提示信息的类型）
    
    RemindType   int64 `json:"remind_type,omitempty"`
    

    // 是否存在超时。可选值:true(是),false(否)
    
    ExistTimeout   bool `json:"exist_timeout,omitempty"`
    

    // 超时时间。格式:yyyy-MM-dd HH:mm:ss
    
    Timeout   string `json:"timeout,omitempty"`
    

}
*/

// RefundRemindTimeout 
type RefundRemindTimeout struct {

    // 提醒的类型（退款详情中提示信息的类型）
    RemindType   int64 `json:"remind_type,omitempty"`

    // 是否存在超时。可选值:true(是),false(否)
    ExistTimeout   bool `json:"exist_timeout,omitempty"`

    // 超时时间。格式:yyyy-MM-dd HH:mm:ss
    Timeout   string `json:"timeout,omitempty"`

}
