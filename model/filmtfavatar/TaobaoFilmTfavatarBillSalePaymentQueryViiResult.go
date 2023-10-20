package filmtfavatar

import (
	"sync"
)

// TaobaoFilmTfavatarBillSalePaymentQueryViiResult 结构体
type TaobaoFilmTfavatarBillSalePaymentQueryViiResult struct {
	// 错误信息
	ReturnMessage string `json:"return_message,omitempty" xml:"return_message,omitempty"`
	// 错误码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 请求ID(排查问题用)
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	ReturnValue *ReturnValue `json:"return_value,omitempty" xml:"return_value,omitempty"`
}

var poolTaobaoFilmTfavatarBillSalePaymentQueryViiResult = sync.Pool{
	New: func() any {
		return new(TaobaoFilmTfavatarBillSalePaymentQueryViiResult)
	},
}

// GetTaobaoFilmTfavatarBillSalePaymentQueryViiResult() 从对象池中获取TaobaoFilmTfavatarBillSalePaymentQueryViiResult
func GetTaobaoFilmTfavatarBillSalePaymentQueryViiResult() *TaobaoFilmTfavatarBillSalePaymentQueryViiResult {
	return poolTaobaoFilmTfavatarBillSalePaymentQueryViiResult.Get().(*TaobaoFilmTfavatarBillSalePaymentQueryViiResult)
}

// ReleaseTaobaoFilmTfavatarBillSalePaymentQueryViiResult 释放TaobaoFilmTfavatarBillSalePaymentQueryViiResult
func ReleaseTaobaoFilmTfavatarBillSalePaymentQueryViiResult(v *TaobaoFilmTfavatarBillSalePaymentQueryViiResult) {
	v.ReturnMessage = ""
	v.ReturnCode = ""
	v.RequestId = ""
	v.ReturnValue = nil
	poolTaobaoFilmTfavatarBillSalePaymentQueryViiResult.Put(v)
}
