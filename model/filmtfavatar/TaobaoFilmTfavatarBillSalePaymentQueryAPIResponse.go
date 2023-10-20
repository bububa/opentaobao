package filmtfavatar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFilmTfavatarBillSalePaymentQueryAPIResponse 获取影院卖品账单--支付账单 API返回值
// taobao.film.tfavatar.bill.sale.payment.query
//
// 获取影院卖品账单--支付账单
type TaobaoFilmTfavatarBillSalePaymentQueryAPIResponse struct {
	model.CommonResponse
	TaobaoFilmTfavatarBillSalePaymentQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFilmTfavatarBillSalePaymentQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFilmTfavatarBillSalePaymentQueryAPIResponseModel).Reset()
}

// TaobaoFilmTfavatarBillSalePaymentQueryAPIResponseModel is 获取影院卖品账单--支付账单 成功返回结果
type TaobaoFilmTfavatarBillSalePaymentQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"film_tfavatar_bill_sale_payment_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoFilmTfavatarBillSalePaymentQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFilmTfavatarBillSalePaymentQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFilmTfavatarBillSalePaymentQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFilmTfavatarBillSalePaymentQueryAPIResponse)
	},
}

// GetTaobaoFilmTfavatarBillSalePaymentQueryAPIResponse 从 sync.Pool 获取 TaobaoFilmTfavatarBillSalePaymentQueryAPIResponse
func GetTaobaoFilmTfavatarBillSalePaymentQueryAPIResponse() *TaobaoFilmTfavatarBillSalePaymentQueryAPIResponse {
	return poolTaobaoFilmTfavatarBillSalePaymentQueryAPIResponse.Get().(*TaobaoFilmTfavatarBillSalePaymentQueryAPIResponse)
}

// ReleaseTaobaoFilmTfavatarBillSalePaymentQueryAPIResponse 将 TaobaoFilmTfavatarBillSalePaymentQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFilmTfavatarBillSalePaymentQueryAPIResponse(v *TaobaoFilmTfavatarBillSalePaymentQueryAPIResponse) {
	v.Reset()
	poolTaobaoFilmTfavatarBillSalePaymentQueryAPIResponse.Put(v)
}
