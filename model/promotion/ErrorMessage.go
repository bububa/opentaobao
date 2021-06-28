package promotion

// ErrorMessage 
/* model for simplify = false
type ErrorMessage struct {

    // 买家昵称
    
    BuyerNick   string `json:"buyer_nick,omitempty"`
    

    // 发送失败的原因
    
    Reason   string `json:"reason,omitempty"`
    

}
*/

// ErrorMessage 
type ErrorMessage struct {

    // 买家昵称
    BuyerNick   string `json:"buyer_nick,omitempty"`

    // 发送失败的原因
    Reason   string `json:"reason,omitempty"`

}
