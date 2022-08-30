package filmtfavatar

import (
	"encoding/xml"

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

// TaobaoFilmTfavatarBillSaleRefundQueryViiAPIResponseModel is 获取影院卖品账单--退款账单-V2版本(正逆分离) 成功返回结果
type TaobaoFilmTfavatarBillSaleRefundQueryViiAPIResponseModel struct {
	XMLName xml.Name `xml:"film_tfavatar_bill_sale_refund_query_vii_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装
	Result *TaobaoFilmTfavatarBillSaleRefundQueryViiResult `json:"result,omitempty" xml:"result,omitempty"`
}
