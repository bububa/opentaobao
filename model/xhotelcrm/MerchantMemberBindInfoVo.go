package xhotelcrm

import (
	"sync"
)

// MerchantMemberBindInfoVo 结构体
type MerchantMemberBindInfoVo struct {
	// 渠道
	Fpt string `json:"fpt,omitempty" xml:"fpt,omitempty"`
	// 注册时间
	Time string `json:"time,omitempty" xml:"time,omitempty"`
	// 1.未注册/非支付宝端注册的会员2.未注册成功3.非新会员
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 是否完成任务
	Status bool `json:"status,omitempty" xml:"status,omitempty"`
}

var poolMerchantMemberBindInfoVo = sync.Pool{
	New: func() any {
		return new(MerchantMemberBindInfoVo)
	},
}

// GetMerchantMemberBindInfoVo() 从对象池中获取MerchantMemberBindInfoVo
func GetMerchantMemberBindInfoVo() *MerchantMemberBindInfoVo {
	return poolMerchantMemberBindInfoVo.Get().(*MerchantMemberBindInfoVo)
}

// ReleaseMerchantMemberBindInfoVo 释放MerchantMemberBindInfoVo
func ReleaseMerchantMemberBindInfoVo(v *MerchantMemberBindInfoVo) {
	v.Fpt = ""
	v.Time = ""
	v.Reason = ""
	v.Status = false
	poolMerchantMemberBindInfoVo.Put(v)
}
