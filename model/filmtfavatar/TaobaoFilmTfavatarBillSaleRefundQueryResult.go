package filmtfavatar

import (
	"sync"
)

// TaobaoFilmTfavatarBillSaleRefundQueryResult 结构体
type TaobaoFilmTfavatarBillSaleRefundQueryResult struct {
	// 返回码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 请求id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msg
	ReturnMessage string `json:"return_message,omitempty" xml:"return_message,omitempty"`
	// 返回参数
	ReturnValue *ReturnValue `json:"return_value,omitempty" xml:"return_value,omitempty"`
}

var poolTaobaoFilmTfavatarBillSaleRefundQueryResult = sync.Pool{
	New: func() any {
		return new(TaobaoFilmTfavatarBillSaleRefundQueryResult)
	},
}

// GetTaobaoFilmTfavatarBillSaleRefundQueryResult() 从对象池中获取TaobaoFilmTfavatarBillSaleRefundQueryResult
func GetTaobaoFilmTfavatarBillSaleRefundQueryResult() *TaobaoFilmTfavatarBillSaleRefundQueryResult {
	return poolTaobaoFilmTfavatarBillSaleRefundQueryResult.Get().(*TaobaoFilmTfavatarBillSaleRefundQueryResult)
}

// ReleaseTaobaoFilmTfavatarBillSaleRefundQueryResult 释放TaobaoFilmTfavatarBillSaleRefundQueryResult
func ReleaseTaobaoFilmTfavatarBillSaleRefundQueryResult(v *TaobaoFilmTfavatarBillSaleRefundQueryResult) {
	v.ReturnCode = ""
	v.RequestId = ""
	v.ReturnMessage = ""
	v.ReturnValue = nil
	poolTaobaoFilmTfavatarBillSaleRefundQueryResult.Put(v)
}
