package filmtfavatar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFilmTfavatarBillTicketRefundQueryAPIResponse 获取影院票务账单-退款账单 API返回值
// taobao.film.tfavatar.bill.ticket.refund.query
//
// 获取影院票务账单-支付订单
// data字段为加密字段, 不可分拆.
type TaobaoFilmTfavatarBillTicketRefundQueryAPIResponse struct {
	model.CommonResponse
	TaobaoFilmTfavatarBillTicketRefundQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFilmTfavatarBillTicketRefundQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFilmTfavatarBillTicketRefundQueryAPIResponseModel).Reset()
}

// TaobaoFilmTfavatarBillTicketRefundQueryAPIResponseModel is 获取影院票务账单-退款账单 成功返回结果
type TaobaoFilmTfavatarBillTicketRefundQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"film_tfavatar_bill_ticket_refund_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoFilmTfavatarBillTicketRefundQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFilmTfavatarBillTicketRefundQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFilmTfavatarBillTicketRefundQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFilmTfavatarBillTicketRefundQueryAPIResponse)
	},
}

// GetTaobaoFilmTfavatarBillTicketRefundQueryAPIResponse 从 sync.Pool 获取 TaobaoFilmTfavatarBillTicketRefundQueryAPIResponse
func GetTaobaoFilmTfavatarBillTicketRefundQueryAPIResponse() *TaobaoFilmTfavatarBillTicketRefundQueryAPIResponse {
	return poolTaobaoFilmTfavatarBillTicketRefundQueryAPIResponse.Get().(*TaobaoFilmTfavatarBillTicketRefundQueryAPIResponse)
}

// ReleaseTaobaoFilmTfavatarBillTicketRefundQueryAPIResponse 将 TaobaoFilmTfavatarBillTicketRefundQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFilmTfavatarBillTicketRefundQueryAPIResponse(v *TaobaoFilmTfavatarBillTicketRefundQueryAPIResponse) {
	v.Reset()
	poolTaobaoFilmTfavatarBillTicketRefundQueryAPIResponse.Put(v)
}
