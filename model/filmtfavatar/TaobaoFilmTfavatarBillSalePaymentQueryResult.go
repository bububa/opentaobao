package filmtfavatar

import (
	"sync"
)

// TaobaoFilmTfavatarBillSalePaymentQueryResult 结构体
type TaobaoFilmTfavatarBillSalePaymentQueryResult struct {
	// 返回码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 请求id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msg
	ReturnMessage string `json:"return_message,omitempty" xml:"return_message,omitempty"`
	// 返回参数
	ReturnValue *ReturnValue `json:"return_value,omitempty" xml:"return_value,omitempty"`
}

var poolTaobaoFilmTfavatarBillSalePaymentQueryResult = sync.Pool{
	New: func() any {
		return new(TaobaoFilmTfavatarBillSalePaymentQueryResult)
	},
}

// GetTaobaoFilmTfavatarBillSalePaymentQueryResult() 从对象池中获取TaobaoFilmTfavatarBillSalePaymentQueryResult
func GetTaobaoFilmTfavatarBillSalePaymentQueryResult() *TaobaoFilmTfavatarBillSalePaymentQueryResult {
	return poolTaobaoFilmTfavatarBillSalePaymentQueryResult.Get().(*TaobaoFilmTfavatarBillSalePaymentQueryResult)
}

// ReleaseTaobaoFilmTfavatarBillSalePaymentQueryResult 释放TaobaoFilmTfavatarBillSalePaymentQueryResult
func ReleaseTaobaoFilmTfavatarBillSalePaymentQueryResult(v *TaobaoFilmTfavatarBillSalePaymentQueryResult) {
	v.ReturnCode = ""
	v.RequestId = ""
	v.ReturnMessage = ""
	v.ReturnValue = nil
	poolTaobaoFilmTfavatarBillSalePaymentQueryResult.Put(v)
}
