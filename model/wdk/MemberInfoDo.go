package wdk

import (
	"sync"
)

// MemberInfoDo 结构体
type MemberInfoDo struct {
	// 会员卡号
	CardNum string `json:"card_num,omitempty" xml:"card_num,omitempty"`
	// 会员卡状态  &#39;使用中&#39; or &#39;已挂失&#39; or &#39;已作废&#39; or &#39;已补发卡&#39; or &#39;已退卡&#39; or &#39;已冻结&#39; or &#39;未激活&#39; or &#39;已坏卡登记&#39; or &#39;已销毁&#39; or &#39;已制卡&#39; or &#39;已发卡&#39; or &#39;已核对&#39; or &#39;已回收&#39; or &#39;空卡&#39; or &#39;异常&#39; or &#39;已损卡&#39;
	State string `json:"state,omitempty" xml:"state,omitempty"`
	// 卡有效期截止日期，格式：yyyy-MM-dd HH:mm:ss
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 扩展属性map
	MemberAttributes string `json:"member_attributes,omitempty" xml:"member_attributes,omitempty"`
	// 会员卡类型
	CardType string `json:"card_type,omitempty" xml:"card_type,omitempty"`
	// 会员卡等级
	CardLevel int64 `json:"card_level,omitempty" xml:"card_level,omitempty"`
	// 如果卡长期有效，值为true，为true时，默认不校验endTime
	Forever bool `json:"forever,omitempty" xml:"forever,omitempty"`
}

var poolMemberInfoDo = sync.Pool{
	New: func() any {
		return new(MemberInfoDo)
	},
}

// GetMemberInfoDo() 从对象池中获取MemberInfoDo
func GetMemberInfoDo() *MemberInfoDo {
	return poolMemberInfoDo.Get().(*MemberInfoDo)
}

// ReleaseMemberInfoDo 释放MemberInfoDo
func ReleaseMemberInfoDo(v *MemberInfoDo) {
	v.CardNum = ""
	v.State = ""
	v.EndTime = ""
	v.MemberAttributes = ""
	v.CardType = ""
	v.CardLevel = 0
	v.Forever = false
	poolMemberInfoDo.Put(v)
}
