package alimember

import (
	"sync"
)

// ExitMemberDto 结构体
type ExitMemberDto struct {
	// 要退会的商家
	OpenMerchantId string `json:"open_merchant_id,omitempty" xml:"open_merchant_id,omitempty"`
	// 商家账号类型
	UidType string `json:"uid_type,omitempty" xml:"uid_type,omitempty"`
	// isv会员id
	IsvMemberId string `json:"isv_member_id,omitempty" xml:"isv_member_id,omitempty"`
}

var poolExitMemberDto = sync.Pool{
	New: func() any {
		return new(ExitMemberDto)
	},
}

// GetExitMemberDto() 从对象池中获取ExitMemberDto
func GetExitMemberDto() *ExitMemberDto {
	return poolExitMemberDto.Get().(*ExitMemberDto)
}

// ReleaseExitMemberDto 释放ExitMemberDto
func ReleaseExitMemberDto(v *ExitMemberDto) {
	v.OpenMerchantId = ""
	v.UidType = ""
	v.IsvMemberId = ""
	poolExitMemberDto.Put(v)
}
