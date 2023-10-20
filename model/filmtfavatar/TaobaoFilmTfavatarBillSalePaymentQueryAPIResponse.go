package filmtfavatar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofilmtfavatarbillsalepaymentqueryAPIResponse 获取影院卖品账单--支付账单 API返回值
// taobao.film.tfavatar.bill.sale.payment.query
//
// 获取影院卖品账单--支付账单
type TaobaofilmtfavatarbillsalepaymentqueryAPIResponse struct {
	model.CommonResponse
	TaobaofilmtfavatarbillsalepaymentqueryAPIResponseModel
}

// TaobaofilmtfavatarbillsalepaymentqueryAPIResponseModel is 获取影院卖品账单--支付账单 成功返回结果
type TaobaofilmtfavatarbillsalepaymentqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"film_tfavatar_bill_sale_payment_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaofilmtfavatarbillsalepaymentqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
