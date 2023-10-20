package omniorder

import (
	"sync"
)

// TaobaoJstAstrolabeStoreinventoryItemupdateError 结构体
type TaobaoJstAstrolabeStoreinventoryItemupdateError struct {
	// 错误描述
	Descrpition string `json:"descrpition,omitempty" xml:"descrpition,omitempty"`
	// 处理失败的流水号
	FailedBillNum string `json:"failed_bill_num,omitempty" xml:"failed_bill_num,omitempty"`
}

var poolTaobaoJstAstrolabeStoreinventoryItemupdateError = sync.Pool{
	New: func() any {
		return new(TaobaoJstAstrolabeStoreinventoryItemupdateError)
	},
}

// GetTaobaoJstAstrolabeStoreinventoryItemupdateError() 从对象池中获取TaobaoJstAstrolabeStoreinventoryItemupdateError
func GetTaobaoJstAstrolabeStoreinventoryItemupdateError() *TaobaoJstAstrolabeStoreinventoryItemupdateError {
	return poolTaobaoJstAstrolabeStoreinventoryItemupdateError.Get().(*TaobaoJstAstrolabeStoreinventoryItemupdateError)
}

// ReleaseTaobaoJstAstrolabeStoreinventoryItemupdateError 释放TaobaoJstAstrolabeStoreinventoryItemupdateError
func ReleaseTaobaoJstAstrolabeStoreinventoryItemupdateError(v *TaobaoJstAstrolabeStoreinventoryItemupdateError) {
	v.Descrpition = ""
	v.FailedBillNum = ""
	poolTaobaoJstAstrolabeStoreinventoryItemupdateError.Put(v)
}
