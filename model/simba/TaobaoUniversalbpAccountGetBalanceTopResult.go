package simba

import (
	"sync"
)

// TaobaoUniversalbpAccountGetBalanceTopResult 结构体
type TaobaoUniversalbpAccountGetBalanceTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	TopAccountBalanceVO *TopAccountBalanceVo `json:"top_account_balance_v_o,omitempty" xml:"top_account_balance_v_o,omitempty"`
}

var poolTaobaoUniversalbpAccountGetBalanceTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpAccountGetBalanceTopResult)
	},
}

// GetTaobaoUniversalbpAccountGetBalanceTopResult() 从对象池中获取TaobaoUniversalbpAccountGetBalanceTopResult
func GetTaobaoUniversalbpAccountGetBalanceTopResult() *TaobaoUniversalbpAccountGetBalanceTopResult {
	return poolTaobaoUniversalbpAccountGetBalanceTopResult.Get().(*TaobaoUniversalbpAccountGetBalanceTopResult)
}

// ReleaseTaobaoUniversalbpAccountGetBalanceTopResult 释放TaobaoUniversalbpAccountGetBalanceTopResult
func ReleaseTaobaoUniversalbpAccountGetBalanceTopResult(v *TaobaoUniversalbpAccountGetBalanceTopResult) {
	v.Info = nil
	v.TopAccountBalanceVO = nil
	poolTaobaoUniversalbpAccountGetBalanceTopResult.Put(v)
}
