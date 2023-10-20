package nrt

import (
	"sync"
)

// MemberSynResponse 结构体
type MemberSynResponse struct {
	// 加密后的taoId
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 业务code
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 外部会员ID
	OutMemberId string `json:"out_member_id,omitempty" xml:"out_member_id,omitempty"`
	// sellerid
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
}

var poolMemberSynResponse = sync.Pool{
	New: func() any {
		return new(MemberSynResponse)
	},
}

// GetMemberSynResponse() 从对象池中获取MemberSynResponse
func GetMemberSynResponse() *MemberSynResponse {
	return poolMemberSynResponse.Get().(*MemberSynResponse)
}

// ReleaseMemberSynResponse 释放MemberSynResponse
func ReleaseMemberSynResponse(v *MemberSynResponse) {
	v.OpenId = ""
	v.BizCode = ""
	v.OutMemberId = ""
	v.SellerId = 0
	poolMemberSynResponse.Put(v)
}
