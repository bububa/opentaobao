package idleisv

import (
	"sync"
)

// IdleDepositTopSubVo 结构体
type IdleDepositTopSubVo struct {
	// 保证金余额是否足够
	BalanceEnough bool `json:"balance_enough,omitempty" xml:"balance_enough,omitempty"`
}

var poolIdleDepositTopSubVo = sync.Pool{
	New: func() any {
		return new(IdleDepositTopSubVo)
	},
}

// GetIdleDepositTopSubVo() 从对象池中获取IdleDepositTopSubVo
func GetIdleDepositTopSubVo() *IdleDepositTopSubVo {
	return poolIdleDepositTopSubVo.Get().(*IdleDepositTopSubVo)
}

// ReleaseIdleDepositTopSubVo 释放IdleDepositTopSubVo
func ReleaseIdleDepositTopSubVo(v *IdleDepositTopSubVo) {
	v.BalanceEnough = false
	poolIdleDepositTopSubVo.Put(v)
}
