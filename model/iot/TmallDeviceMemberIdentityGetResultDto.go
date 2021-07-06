package iot

// TmallDeviceMemberIdentityGetResultDto 结构体
type TmallDeviceMemberIdentityGetResultDto struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// total
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// result
	MemberInfo *MemberAccountDto `json:"member_info,omitempty" xml:"member_info,omitempty"`
}
