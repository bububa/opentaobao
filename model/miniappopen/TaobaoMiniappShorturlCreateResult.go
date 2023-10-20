package miniappopen

import (
	"sync"
)

// TaobaoMiniappShorturlCreateResult 结构体
type TaobaoMiniappShorturlCreateResult struct {
	// model
	Model *MiniAppShortUrlDto `json:"model,omitempty" xml:"model,omitempty"`
}

var poolTaobaoMiniappShorturlCreateResult = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappShorturlCreateResult)
	},
}

// GetTaobaoMiniappShorturlCreateResult() 从对象池中获取TaobaoMiniappShorturlCreateResult
func GetTaobaoMiniappShorturlCreateResult() *TaobaoMiniappShorturlCreateResult {
	return poolTaobaoMiniappShorturlCreateResult.Get().(*TaobaoMiniappShorturlCreateResult)
}

// ReleaseTaobaoMiniappShorturlCreateResult 释放TaobaoMiniappShorturlCreateResult
func ReleaseTaobaoMiniappShorturlCreateResult(v *TaobaoMiniappShorturlCreateResult) {
	v.Model = nil
	poolTaobaoMiniappShorturlCreateResult.Put(v)
}
