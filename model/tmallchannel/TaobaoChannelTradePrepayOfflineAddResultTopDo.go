package tmallchannel

import (
	"sync"
)

// TaobaoChannelTradePrepayOfflineAddResultTopDo 结构体
type TaobaoChannelTradePrepayOfflineAddResultTopDo struct {
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoChannelTradePrepayOfflineAddResultTopDo = sync.Pool{
	New: func() any {
		return new(TaobaoChannelTradePrepayOfflineAddResultTopDo)
	},
}

// GetTaobaoChannelTradePrepayOfflineAddResultTopDo() 从对象池中获取TaobaoChannelTradePrepayOfflineAddResultTopDo
func GetTaobaoChannelTradePrepayOfflineAddResultTopDo() *TaobaoChannelTradePrepayOfflineAddResultTopDo {
	return poolTaobaoChannelTradePrepayOfflineAddResultTopDo.Get().(*TaobaoChannelTradePrepayOfflineAddResultTopDo)
}

// ReleaseTaobaoChannelTradePrepayOfflineAddResultTopDo 释放TaobaoChannelTradePrepayOfflineAddResultTopDo
func ReleaseTaobaoChannelTradePrepayOfflineAddResultTopDo(v *TaobaoChannelTradePrepayOfflineAddResultTopDo) {
	v.Success = false
	poolTaobaoChannelTradePrepayOfflineAddResultTopDo.Put(v)
}
