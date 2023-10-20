package traderate

import (
	"sync"
)

// TaobaoFliggyWrateGetmixratelistResult 结构体
type TaobaoFliggyWrateGetmixratelistResult struct {
	// 返回对象
	Model *GetMixRateListResult `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFliggyWrateGetmixratelistResult = sync.Pool{
	New: func() any {
		return new(TaobaoFliggyWrateGetmixratelistResult)
	},
}

// GetTaobaoFliggyWrateGetmixratelistResult() 从对象池中获取TaobaoFliggyWrateGetmixratelistResult
func GetTaobaoFliggyWrateGetmixratelistResult() *TaobaoFliggyWrateGetmixratelistResult {
	return poolTaobaoFliggyWrateGetmixratelistResult.Get().(*TaobaoFliggyWrateGetmixratelistResult)
}

// ReleaseTaobaoFliggyWrateGetmixratelistResult 释放TaobaoFliggyWrateGetmixratelistResult
func ReleaseTaobaoFliggyWrateGetmixratelistResult(v *TaobaoFliggyWrateGetmixratelistResult) {
	v.Model = nil
	v.Success = false
	poolTaobaoFliggyWrateGetmixratelistResult.Put(v)
}
