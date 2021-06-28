package iot

// TmallDeviceMemberIdentityGetResultDto 
/* model for simplify = false
type TmallDeviceMemberIdentityGetResultDto struct {

    // total
    
    Total   int64 `json:"total,omitempty"`
    

    // result
    
    MemberInfo  *struct {
        MemberAccountDto  *MemberAccountDto `json:"member_account_dto,omitempty"`
    } `json:"member_info,omitempty"`
    

    // code
    
    Code   string `json:"code,omitempty"`
    

    // msg
    
    Msg   string `json:"msg,omitempty"`
    

}
*/

// TmallDeviceMemberIdentityGetResultDto 
type TmallDeviceMemberIdentityGetResultDto struct {

    // total
    Total   int64 `json:"total,omitempty"`

    // result
    MemberInfo   *MemberAccountDto `json:"member_info,omitempty"`

    // code
    Code   string `json:"code,omitempty"`

    // msg
    Msg   string `json:"msg,omitempty"`

}
