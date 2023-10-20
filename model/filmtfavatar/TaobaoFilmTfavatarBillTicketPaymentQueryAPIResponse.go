package filmtfavatar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFilmTfavatarBillTicketPaymentQueryAPIResponse 获取影院票务账单-支付订单 API返回值
// taobao.film.tfavatar.bill.ticket.payment.query
//
// 获取影院票务账单-支付订单
type TaobaoFilmTfavatarBillTicketPaymentQueryAPIResponse struct {
	model.CommonResponse
	TaobaoFilmTfavatarBillTicketPaymentQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFilmTfavatarBillTicketPaymentQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFilmTfavatarBillTicketPaymentQueryAPIResponseModel).Reset()
}

// TaobaoFilmTfavatarBillTicketPaymentQueryAPIResponseModel is 获取影院票务账单-支付订单 成功返回结果
type TaobaoFilmTfavatarBillTicketPaymentQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"film_tfavatar_bill_ticket_payment_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoFilmTfavatarBillTicketPaymentQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFilmTfavatarBillTicketPaymentQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFilmTfavatarBillTicketPaymentQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFilmTfavatarBillTicketPaymentQueryAPIResponse)
	},
}

// GetTaobaoFilmTfavatarBillTicketPaymentQueryAPIResponse 从 sync.Pool 获取 TaobaoFilmTfavatarBillTicketPaymentQueryAPIResponse
func GetTaobaoFilmTfavatarBillTicketPaymentQueryAPIResponse() *TaobaoFilmTfavatarBillTicketPaymentQueryAPIResponse {
	return poolTaobaoFilmTfavatarBillTicketPaymentQueryAPIResponse.Get().(*TaobaoFilmTfavatarBillTicketPaymentQueryAPIResponse)
}

// ReleaseTaobaoFilmTfavatarBillTicketPaymentQueryAPIResponse 将 TaobaoFilmTfavatarBillTicketPaymentQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFilmTfavatarBillTicketPaymentQueryAPIResponse(v *TaobaoFilmTfavatarBillTicketPaymentQueryAPIResponse) {
	v.Reset()
	poolTaobaoFilmTfavatarBillTicketPaymentQueryAPIResponse.Put(v)
}
