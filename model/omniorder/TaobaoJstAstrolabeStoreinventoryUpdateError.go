package omniorder

import (
	"sync"
)

// TaobaoJstAstrolabeStoreinventoryUpdateError 结构体
type TaobaoJstAstrolabeStoreinventoryUpdateError struct {
	// 错误描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 处理失败的流水号
	FailedBillNum string `json:"failed_bill_num,omitempty" xml:"failed_bill_num,omitempty"`
}

var poolTaobaoJstAstrolabeStoreinventoryUpdateError = sync.Pool{
	New: func() any {
		return new(TaobaoJstAstrolabeStoreinventoryUpdateError)
	},
}

// GetTaobaoJstAstrolabeStoreinventoryUpdateError() 从对象池中获取TaobaoJstAstrolabeStoreinventoryUpdateError
func GetTaobaoJstAstrolabeStoreinventoryUpdateError() *TaobaoJstAstrolabeStoreinventoryUpdateError {
	return poolTaobaoJstAstrolabeStoreinventoryUpdateError.Get().(*TaobaoJstAstrolabeStoreinventoryUpdateError)
}

// ReleaseTaobaoJstAstrolabeStoreinventoryUpdateError 释放TaobaoJstAstrolabeStoreinventoryUpdateError
func ReleaseTaobaoJstAstrolabeStoreinventoryUpdateError(v *TaobaoJstAstrolabeStoreinventoryUpdateError) {
	v.Description = ""
	v.FailedBillNum = ""
	poolTaobaoJstAstrolabeStoreinventoryUpdateError.Put(v)
}
