package fenxiao

import (
	"sync"
)

// TaobaoFenxiaoTradePrepayOfflineAddResultTopDo 结构体
type TaobaoFenxiaoTradePrepayOfflineAddResultTopDo struct {
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFenxiaoTradePrepayOfflineAddResultTopDo = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoTradePrepayOfflineAddResultTopDo)
	},
}

// GetTaobaoFenxiaoTradePrepayOfflineAddResultTopDo() 从对象池中获取TaobaoFenxiaoTradePrepayOfflineAddResultTopDo
func GetTaobaoFenxiaoTradePrepayOfflineAddResultTopDo() *TaobaoFenxiaoTradePrepayOfflineAddResultTopDo {
	return poolTaobaoFenxiaoTradePrepayOfflineAddResultTopDo.Get().(*TaobaoFenxiaoTradePrepayOfflineAddResultTopDo)
}

// ReleaseTaobaoFenxiaoTradePrepayOfflineAddResultTopDo 释放TaobaoFenxiaoTradePrepayOfflineAddResultTopDo
func ReleaseTaobaoFenxiaoTradePrepayOfflineAddResultTopDo(v *TaobaoFenxiaoTradePrepayOfflineAddResultTopDo) {
	v.Success = false
	poolTaobaoFenxiaoTradePrepayOfflineAddResultTopDo.Put(v)
}
