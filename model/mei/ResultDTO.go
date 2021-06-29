package mei

// ResultDTO 
type ResultDTO struct {
    // total
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
    // result
    Result   *MemberAccountDTO `json:"result,omitempty" xml:"result,omitempty"`
    // code
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // msg
    Msg   string `json:"msg,omitempty" xml:"msg,omitempty"`
}
