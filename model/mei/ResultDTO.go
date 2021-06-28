package mei

// ResultDTO 
/* model for simplify = false
type ResultDTO struct {

    // total
    
    Total   int64 `json:"total,omitempty"`
    

    // result
    
    Result  *struct {
        MemberAccountDTO  *MemberAccountDTO `json:"member_account_dto,omitempty"`
    } `json:"result,omitempty"`
    

    // code
    
    Code   string `json:"code,omitempty"`
    

    // msg
    
    Msg   string `json:"msg,omitempty"`
    

}
*/

// ResultDTO 
type ResultDTO struct {

    // total
    Total   int64 `json:"total,omitempty"`

    // result
    Result   *MemberAccountDTO `json:"result,omitempty"`

    // code
    Code   string `json:"code,omitempty"`

    // msg
    Msg   string `json:"msg,omitempty"`

}
