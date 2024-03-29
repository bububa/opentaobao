package wdk

import (
	"sync"
)

// OrderPayInfoBO 结构体
type OrderPayInfoBO struct {
	// 买家用户名
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// 买家昵称
	UserNick string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
	// 买家电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 商家会员卡号
	MemberCardNum string `json:"member_card_num,omitempty" xml:"member_card_num,omitempty"`
	// 会员标
	MemberTags string `json:"member_tags,omitempty" xml:"member_tags,omitempty"`
}

var poolOrderPayInfoBO = sync.Pool{
	New: func() any {
		return new(OrderPayInfoBO)
	},
}

// GetOrderPayInfoBO() 从对象池中获取OrderPayInfoBO
func GetOrderPayInfoBO() *OrderPayInfoBO {
	return poolOrderPayInfoBO.Get().(*OrderPayInfoBO)
}

// ReleaseOrderPayInfoBO 释放OrderPayInfoBO
func ReleaseOrderPayInfoBO(v *OrderPayInfoBO) {
	v.UserName = ""
	v.UserNick = ""
	v.Phone = ""
	v.MemberCardNum = ""
	v.MemberTags = ""
	poolOrderPayInfoBO.Put(v)
}
