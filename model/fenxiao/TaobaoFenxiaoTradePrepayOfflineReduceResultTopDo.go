package fenxiao

import (
	"sync"
)

// TaobaoFenxiaoTradePrepayOfflineReduceResultTopDo 结构体
type TaobaoFenxiaoTradePrepayOfflineReduceResultTopDo struct {
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFenxiaoTradePrepayOfflineReduceResultTopDo = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoTradePrepayOfflineReduceResultTopDo)
	},
}

// GetTaobaoFenxiaoTradePrepayOfflineReduceResultTopDo() 从对象池中获取TaobaoFenxiaoTradePrepayOfflineReduceResultTopDo
func GetTaobaoFenxiaoTradePrepayOfflineReduceResultTopDo() *TaobaoFenxiaoTradePrepayOfflineReduceResultTopDo {
	return poolTaobaoFenxiaoTradePrepayOfflineReduceResultTopDo.Get().(*TaobaoFenxiaoTradePrepayOfflineReduceResultTopDo)
}

// ReleaseTaobaoFenxiaoTradePrepayOfflineReduceResultTopDo 释放TaobaoFenxiaoTradePrepayOfflineReduceResultTopDo
func ReleaseTaobaoFenxiaoTradePrepayOfflineReduceResultTopDo(v *TaobaoFenxiaoTradePrepayOfflineReduceResultTopDo) {
	v.Success = false
	poolTaobaoFenxiaoTradePrepayOfflineReduceResultTopDo.Put(v)
}
