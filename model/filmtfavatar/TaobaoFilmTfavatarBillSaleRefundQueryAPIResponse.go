package filmtfavatar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFilmTfavatarBillSaleRefundQueryAPIResponse 获取影院卖品账单--退款账单 API返回值
// taobao.film.tfavatar.bill.sale.refund.query
//
// 获取影院卖品账单--退款账单
type TaobaoFilmTfavatarBillSaleRefundQueryAPIResponse struct {
	model.CommonResponse
	TaobaoFilmTfavatarBillSaleRefundQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFilmTfavatarBillSaleRefundQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFilmTfavatarBillSaleRefundQueryAPIResponseModel).Reset()
}

// TaobaoFilmTfavatarBillSaleRefundQueryAPIResponseModel is 获取影院卖品账单--退款账单 成功返回结果
type TaobaoFilmTfavatarBillSaleRefundQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"film_tfavatar_bill_sale_refund_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoFilmTfavatarBillSaleRefundQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFilmTfavatarBillSaleRefundQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFilmTfavatarBillSaleRefundQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFilmTfavatarBillSaleRefundQueryAPIResponse)
	},
}

// GetTaobaoFilmTfavatarBillSaleRefundQueryAPIResponse 从 sync.Pool 获取 TaobaoFilmTfavatarBillSaleRefundQueryAPIResponse
func GetTaobaoFilmTfavatarBillSaleRefundQueryAPIResponse() *TaobaoFilmTfavatarBillSaleRefundQueryAPIResponse {
	return poolTaobaoFilmTfavatarBillSaleRefundQueryAPIResponse.Get().(*TaobaoFilmTfavatarBillSaleRefundQueryAPIResponse)
}

// ReleaseTaobaoFilmTfavatarBillSaleRefundQueryAPIResponse 将 TaobaoFilmTfavatarBillSaleRefundQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFilmTfavatarBillSaleRefundQueryAPIResponse(v *TaobaoFilmTfavatarBillSaleRefundQueryAPIResponse) {
	v.Reset()
	poolTaobaoFilmTfavatarBillSaleRefundQueryAPIResponse.Put(v)
}
