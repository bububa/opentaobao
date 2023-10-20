package alitripmerchant

import (
	"sync"
)

// MemberCardParamVo 结构体
type MemberCardParamVo struct {
	// 微信公众号会员卡模板 id
	CardId string `json:"card_id,omitempty" xml:"card_id,omitempty"`
	// 用户会员卡 卡号
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 当前用户在微信公众号下的OpenId
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 当前时间戳 (秒)
	Timestamp string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// 加密签名
	Signature string `json:"signature,omitempty" xml:"signature,omitempty"`
	// 加密随机字符串
	NonceStr string `json:"nonce_str,omitempty" xml:"nonce_str,omitempty"`
}

var poolMemberCardParamVo = sync.Pool{
	New: func() any {
		return new(MemberCardParamVo)
	},
}

// GetMemberCardParamVo() 从对象池中获取MemberCardParamVo
func GetMemberCardParamVo() *MemberCardParamVo {
	return poolMemberCardParamVo.Get().(*MemberCardParamVo)
}

// ReleaseMemberCardParamVo 释放MemberCardParamVo
func ReleaseMemberCardParamVo(v *MemberCardParamVo) {
	v.CardId = ""
	v.Code = ""
	v.OpenId = ""
	v.Timestamp = ""
	v.Signature = ""
	v.NonceStr = ""
	poolMemberCardParamVo.Put(v)
}
