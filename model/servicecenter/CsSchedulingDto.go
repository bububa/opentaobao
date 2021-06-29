package servicecenter

// CsSchedulingDTO 
type CsSchedulingDTO struct {
    // 排班起始时间
    StartTime   string `json:"start_time,omitempty" xml:"start_time,omitempty"`
    // 排班结束时间
    EndTime   string `json:"end_time,omitempty" xml:"end_time,omitempty"`
    // 商家子账号昵称
    SellerNick   string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
    // 服务商子账号昵称
    SpNick   string `json:"sp_nick,omitempty" xml:"sp_nick,omitempty"`
}
