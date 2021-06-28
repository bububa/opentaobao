package promotion

// ErrorMessage 
type ErrorMessage struct {

    // 买家昵称
    
    BuyerNick   string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
    

    // 发送失败的原因
    
    Reason   string `json:"reason,omitempty" xml:"reason,omitempty"`
    

}
