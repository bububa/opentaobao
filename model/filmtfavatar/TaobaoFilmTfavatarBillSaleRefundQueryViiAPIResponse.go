package filmtfavatar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFilmTfavatarBillSaleRefundQueryViiAPIResponse 获取影院卖品账单--退款账单-V2版本(正逆分离) API返回值
// taobao.film.tfavatar.bill.sale.refund.query.vii
//
// 获取影院卖品账单--退款账单-V2版本(正逆分离)
type TaobaoFilmTfavatarBillSaleRefundQueryViiAPIResponse struct {
	model.CommonResponse
	TaobaoFilmTfavatarBillSaleRefundQueryViiAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFilmTfavatarBillSaleRefundQueryViiAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFilmTfavatarBillSaleRefundQueryViiAPIResponseModel).Reset()
}

// TaobaoFilmTfavatarBillSaleRefundQueryViiAPIResponseModel is 获取影院卖品账单--退款账单-V2版本(正逆分离) 成功返回结果
type TaobaoFilmTfavatarBillSaleRefundQueryViiAPIResponseModel struct {
	XMLName xml.Name `xml:"film_tfavatar_bill_sale_refund_query_vii_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装
	Result *TaobaoFilmTfavatarBillSaleRefundQueryViiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFilmTfavatarBillSaleRefundQueryViiAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFilmTfavatarBillSaleRefundQueryViiAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFilmTfavatarBillSaleRefundQueryViiAPIResponse)
	},
}

// GetTaobaoFilmTfavatarBillSaleRefundQueryViiAPIResponse 从 sync.Pool 获取 TaobaoFilmTfavatarBillSaleRefundQueryViiAPIResponse
func GetTaobaoFilmTfavatarBillSaleRefundQueryViiAPIResponse() *TaobaoFilmTfavatarBillSaleRefundQueryViiAPIResponse {
	return poolTaobaoFilmTfavatarBillSaleRefundQueryViiAPIResponse.Get().(*TaobaoFilmTfavatarBillSaleRefundQueryViiAPIResponse)
}

// ReleaseTaobaoFilmTfavatarBillSaleRefundQueryViiAPIResponse 将 TaobaoFilmTfavatarBillSaleRefundQueryViiAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFilmTfavatarBillSaleRefundQueryViiAPIResponse(v *TaobaoFilmTfavatarBillSaleRefundQueryViiAPIResponse) {
	v.Reset()
	poolTaobaoFilmTfavatarBillSaleRefundQueryViiAPIResponse.Put(v)
}
