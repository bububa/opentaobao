package mei

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
