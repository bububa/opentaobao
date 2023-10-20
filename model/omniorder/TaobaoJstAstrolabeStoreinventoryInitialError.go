package omniorder

import (
	"sync"
)

// TaobaoJstAstrolabeStoreinventoryInitialError 结构体
type TaobaoJstAstrolabeStoreinventoryInitialError struct {
	// 错误描述
	Descrpition string `json:"descrpition,omitempty" xml:"descrpition,omitempty"`
	// 处理失败的流水号（有多个时，用逗号分隔）
	FailedBillNum string `json:"failed_bill_num,omitempty" xml:"failed_bill_num,omitempty"`
}

var poolTaobaoJstAstrolabeStoreinventoryInitialError = sync.Pool{
	New: func() any {
		return new(TaobaoJstAstrolabeStoreinventoryInitialError)
	},
}

// GetTaobaoJstAstrolabeStoreinventoryInitialError() 从对象池中获取TaobaoJstAstrolabeStoreinventoryInitialError
func GetTaobaoJstAstrolabeStoreinventoryInitialError() *TaobaoJstAstrolabeStoreinventoryInitialError {
	return poolTaobaoJstAstrolabeStoreinventoryInitialError.Get().(*TaobaoJstAstrolabeStoreinventoryInitialError)
}

// ReleaseTaobaoJstAstrolabeStoreinventoryInitialError 释放TaobaoJstAstrolabeStoreinventoryInitialError
func ReleaseTaobaoJstAstrolabeStoreinventoryInitialError(v *TaobaoJstAstrolabeStoreinventoryInitialError) {
	v.Descrpition = ""
	v.FailedBillNum = ""
	poolTaobaoJstAstrolabeStoreinventoryInitialError.Put(v)
}
