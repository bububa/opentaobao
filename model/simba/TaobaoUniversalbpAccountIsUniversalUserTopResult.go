package simba

import (
	"sync"
)

// TaobaoUniversalbpAccountIsUniversalUserTopResult 结构体
type TaobaoUniversalbpAccountIsUniversalUserTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	TopAccountStatusVO *TopAccountStatusVo `json:"top_account_status_v_o,omitempty" xml:"top_account_status_v_o,omitempty"`
}

var poolTaobaoUniversalbpAccountIsUniversalUserTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpAccountIsUniversalUserTopResult)
	},
}

// GetTaobaoUniversalbpAccountIsUniversalUserTopResult() 从对象池中获取TaobaoUniversalbpAccountIsUniversalUserTopResult
func GetTaobaoUniversalbpAccountIsUniversalUserTopResult() *TaobaoUniversalbpAccountIsUniversalUserTopResult {
	return poolTaobaoUniversalbpAccountIsUniversalUserTopResult.Get().(*TaobaoUniversalbpAccountIsUniversalUserTopResult)
}

// ReleaseTaobaoUniversalbpAccountIsUniversalUserTopResult 释放TaobaoUniversalbpAccountIsUniversalUserTopResult
func ReleaseTaobaoUniversalbpAccountIsUniversalUserTopResult(v *TaobaoUniversalbpAccountIsUniversalUserTopResult) {
	v.Info = nil
	v.TopAccountStatusVO = nil
	poolTaobaoUniversalbpAccountIsUniversalUserTopResult.Put(v)
}
