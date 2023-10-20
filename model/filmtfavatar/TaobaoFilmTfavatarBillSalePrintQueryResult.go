package filmtfavatar

import (
	"sync"
)

// TaobaoFilmTfavatarBillSalePrintQueryResult 结构体
type TaobaoFilmTfavatarBillSalePrintQueryResult struct {
	// 返回码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 请求id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msg
	ReturnMessage string `json:"return_message,omitempty" xml:"return_message,omitempty"`
	// 返回参数
	ReturnValue *ReturnValue `json:"return_value,omitempty" xml:"return_value,omitempty"`
}

var poolTaobaoFilmTfavatarBillSalePrintQueryResult = sync.Pool{
	New: func() any {
		return new(TaobaoFilmTfavatarBillSalePrintQueryResult)
	},
}

// GetTaobaoFilmTfavatarBillSalePrintQueryResult() 从对象池中获取TaobaoFilmTfavatarBillSalePrintQueryResult
func GetTaobaoFilmTfavatarBillSalePrintQueryResult() *TaobaoFilmTfavatarBillSalePrintQueryResult {
	return poolTaobaoFilmTfavatarBillSalePrintQueryResult.Get().(*TaobaoFilmTfavatarBillSalePrintQueryResult)
}

// ReleaseTaobaoFilmTfavatarBillSalePrintQueryResult 释放TaobaoFilmTfavatarBillSalePrintQueryResult
func ReleaseTaobaoFilmTfavatarBillSalePrintQueryResult(v *TaobaoFilmTfavatarBillSalePrintQueryResult) {
	v.ReturnCode = ""
	v.RequestId = ""
	v.ReturnMessage = ""
	v.ReturnValue = nil
	poolTaobaoFilmTfavatarBillSalePrintQueryResult.Put(v)
}
