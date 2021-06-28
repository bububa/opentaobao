package mei

// MemberAccountDTO 
/* model for simplify = false
type MemberAccountDTO struct {

    // mixMobile，只有有权限的才有值
    
    MixMobile   string `json:"mix_mobile,omitempty"`
    

    // buyerNick，只有有权限的才有值
    
    BuyerNick   string `json:"buyer_nick,omitempty"`
    

    // 明文手机号，只有有权限的才有值
    
    ClearMobile   string `json:"clear_mobile,omitempty"`
    

}
*/

// MemberAccountDTO 
type MemberAccountDTO struct {

    // mixMobile，只有有权限的才有值
    MixMobile   string `json:"mix_mobile,omitempty"`

    // buyerNick，只有有权限的才有值
    BuyerNick   string `json:"buyer_nick,omitempty"`

    // 明文手机号，只有有权限的才有值
    ClearMobile   string `json:"clear_mobile,omitempty"`

}
