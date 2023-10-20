package tmallchannel

import (
	"sync"
)

// TaobaoChannelTradePrepayOfflineReduceResultTopDo 结构体
type TaobaoChannelTradePrepayOfflineReduceResultTopDo struct {
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoChannelTradePrepayOfflineReduceResultTopDo = sync.Pool{
	New: func() any {
		return new(TaobaoChannelTradePrepayOfflineReduceResultTopDo)
	},
}

// GetTaobaoChannelTradePrepayOfflineReduceResultTopDo() 从对象池中获取TaobaoChannelTradePrepayOfflineReduceResultTopDo
func GetTaobaoChannelTradePrepayOfflineReduceResultTopDo() *TaobaoChannelTradePrepayOfflineReduceResultTopDo {
	return poolTaobaoChannelTradePrepayOfflineReduceResultTopDo.Get().(*TaobaoChannelTradePrepayOfflineReduceResultTopDo)
}

// ReleaseTaobaoChannelTradePrepayOfflineReduceResultTopDo 释放TaobaoChannelTradePrepayOfflineReduceResultTopDo
func ReleaseTaobaoChannelTradePrepayOfflineReduceResultTopDo(v *TaobaoChannelTradePrepayOfflineReduceResultTopDo) {
	v.Success = false
	poolTaobaoChannelTradePrepayOfflineReduceResultTopDo.Put(v)
}
