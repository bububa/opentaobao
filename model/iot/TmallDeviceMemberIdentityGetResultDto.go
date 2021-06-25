package iot

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
