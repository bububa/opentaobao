package iot

// TmallDeviceMemberIdentityGetResultDTO 
type TmallDeviceMemberIdentityGetResultDTO struct {
    // total
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
    // result
    MemberInfo   *MemberAccountDTO `json:"member_info,omitempty" xml:"member_info,omitempty"`
    // code
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // msg
    Msg   string `json:"msg,omitempty" xml:"msg,omitempty"`
}
