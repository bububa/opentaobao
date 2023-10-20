package omniorder

import (
	"sync"
)

// TaobaoJstAstrolabeStoreinventoryIteminitialError 结构体
type TaobaoJstAstrolabeStoreinventoryIteminitialError struct {
	// 错误描述
	Descrpition string `json:"descrpition,omitempty" xml:"descrpition,omitempty"`
	// 处理失败的流水号（有多个时，用逗号分隔）
	FailedBillNum string `json:"failed_bill_num,omitempty" xml:"failed_bill_num,omitempty"`
}

var poolTaobaoJstAstrolabeStoreinventoryIteminitialError = sync.Pool{
	New: func() any {
		return new(TaobaoJstAstrolabeStoreinventoryIteminitialError)
	},
}

// GetTaobaoJstAstrolabeStoreinventoryIteminitialError() 从对象池中获取TaobaoJstAstrolabeStoreinventoryIteminitialError
func GetTaobaoJstAstrolabeStoreinventoryIteminitialError() *TaobaoJstAstrolabeStoreinventoryIteminitialError {
	return poolTaobaoJstAstrolabeStoreinventoryIteminitialError.Get().(*TaobaoJstAstrolabeStoreinventoryIteminitialError)
}

// ReleaseTaobaoJstAstrolabeStoreinventoryIteminitialError 释放TaobaoJstAstrolabeStoreinventoryIteminitialError
func ReleaseTaobaoJstAstrolabeStoreinventoryIteminitialError(v *TaobaoJstAstrolabeStoreinventoryIteminitialError) {
	v.Descrpition = ""
	v.FailedBillNum = ""
	poolTaobaoJstAstrolabeStoreinventoryIteminitialError.Put(v)
}
