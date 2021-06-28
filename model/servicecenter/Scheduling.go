package servicecenter

// Scheduling 
/* model for simplify = false
type Scheduling struct {

    // 排班起始时间
    
    StartTime   string `json:"start_time,omitempty"`
    

    // 排班结束时间
    
    EndTime   string `json:"end_time,omitempty"`
    

    // 商家子账号
    
    SellerNick   string `json:"seller_nick,omitempty"`
    

    // 服务商子账号
    
    SpNick   string `json:"sp_nick,omitempty"`
    

    // 排班记录状态，1表示生效，-1表示失效
    
    State   int64 `json:"state,omitempty"`
    

    // 排班记录状态描述
    
    StateDes   string `json:"state_des,omitempty"`
    

}
*/

// Scheduling 
type Scheduling struct {

    // 排班起始时间
    StartTime   string `json:"start_time,omitempty"`

    // 排班结束时间
    EndTime   string `json:"end_time,omitempty"`

    // 商家子账号
    SellerNick   string `json:"seller_nick,omitempty"`

    // 服务商子账号
    SpNick   string `json:"sp_nick,omitempty"`

    // 排班记录状态，1表示生效，-1表示失效
    State   int64 `json:"state,omitempty"`

    // 排班记录状态描述
    StateDes   string `json:"state_des,omitempty"`

}
