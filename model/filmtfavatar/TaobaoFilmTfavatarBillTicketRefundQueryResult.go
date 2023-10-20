package filmtfavatar

import (
	"sync"
)

// TaobaoFilmTfavatarBillTicketRefundQueryResult 结构体
type TaobaoFilmTfavatarBillTicketRefundQueryResult struct {
	// 返回码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 请求id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msg
	ReturnMessage string `json:"return_message,omitempty" xml:"return_message,omitempty"`
	// 返回参数
	ReturnValue *ReturnValue `json:"return_value,omitempty" xml:"return_value,omitempty"`
}

var poolTaobaoFilmTfavatarBillTicketRefundQueryResult = sync.Pool{
	New: func() any {
		return new(TaobaoFilmTfavatarBillTicketRefundQueryResult)
	},
}

// GetTaobaoFilmTfavatarBillTicketRefundQueryResult() 从对象池中获取TaobaoFilmTfavatarBillTicketRefundQueryResult
func GetTaobaoFilmTfavatarBillTicketRefundQueryResult() *TaobaoFilmTfavatarBillTicketRefundQueryResult {
	return poolTaobaoFilmTfavatarBillTicketRefundQueryResult.Get().(*TaobaoFilmTfavatarBillTicketRefundQueryResult)
}

// ReleaseTaobaoFilmTfavatarBillTicketRefundQueryResult 释放TaobaoFilmTfavatarBillTicketRefundQueryResult
func ReleaseTaobaoFilmTfavatarBillTicketRefundQueryResult(v *TaobaoFilmTfavatarBillTicketRefundQueryResult) {
	v.ReturnCode = ""
	v.RequestId = ""
	v.ReturnMessage = ""
	v.ReturnValue = nil
	poolTaobaoFilmTfavatarBillTicketRefundQueryResult.Put(v)
}
