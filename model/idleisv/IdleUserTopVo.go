package idleisv

import (
	"sync"
)

// IdleUserTopVo 结构体
type IdleUserTopVo struct {
	// 用户选项信息: IDLE_UP（闲鱼升级用户），SDR（已开通七天无理由服务），NFR（已开通描述不符包邮退服务），VNR（已开通虚拟-描述不符包退），FD_24HS（已开通极速发货-24小时），FD_10MS（已开通极速发货-10分钟）
	Options []string `json:"options,omitempty" xml:"options>string,omitempty"`
	// 加密用户id
	EncryptUserid string `json:"encrypt_userid,omitempty" xml:"encrypt_userid,omitempty"`
	// 服务保障金信息
	ServiceDeposit *IdleDepositTopSubVo `json:"service_deposit,omitempty" xml:"service_deposit,omitempty"`
}

var poolIdleUserTopVo = sync.Pool{
	New: func() any {
		return new(IdleUserTopVo)
	},
}

// GetIdleUserTopVo() 从对象池中获取IdleUserTopVo
func GetIdleUserTopVo() *IdleUserTopVo {
	return poolIdleUserTopVo.Get().(*IdleUserTopVo)
}

// ReleaseIdleUserTopVo 释放IdleUserTopVo
func ReleaseIdleUserTopVo(v *IdleUserTopVo) {
	v.Options = v.Options[:0]
	v.EncryptUserid = ""
	v.ServiceDeposit = nil
	poolIdleUserTopVo.Put(v)
}
