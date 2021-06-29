package feedflow

// AccountDTO 
type AccountDTO struct {
    // 账户余额，单位：元
    Balance   string `json:"balance,omitempty" xml:"balance,omitempty"`
    // 现金余额，单位：元
    CashBalance   string `json:"cash_balance,omitempty" xml:"cash_balance,omitempty"`
    // 可用余额，单位：元
    AvailableBalance   string `json:"available_balance,omitempty" xml:"available_balance,omitempty"`
    // 红包，单位：元
    RedPacket   string `json:"red_packet,omitempty" xml:"red_packet,omitempty"`
}
