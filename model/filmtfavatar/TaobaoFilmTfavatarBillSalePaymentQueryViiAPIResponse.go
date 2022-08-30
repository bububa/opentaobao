package filmtfavatar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFilmTfavatarBillSalePaymentQueryViiAPIResponse 获取影院卖品账单--支付账单-V2版本(正逆分离) API返回值
// taobao.film.tfavatar.bill.sale.payment.query.vii
//
// 获取影院卖品账单--支付账单-V2版本(正逆分离)
type TaobaoFilmTfavatarBillSalePaymentQueryViiAPIResponse struct {
	model.CommonResponse
	TaobaoFilmTfavatarBillSalePaymentQueryViiAPIResponseModel
}

// TaobaoFilmTfavatarBillSalePaymentQueryViiAPIResponseModel is 获取影院卖品账单--支付账单-V2版本(正逆分离) 成功返回结果
type TaobaoFilmTfavatarBillSalePaymentQueryViiAPIResponseModel struct {
	XMLName xml.Name `xml:"film_tfavatar_bill_sale_payment_query_vii_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装
	Result *TaobaoFilmTfavatarBillSalePaymentQueryViiResult `json:"result,omitempty" xml:"result,omitempty"`
}
