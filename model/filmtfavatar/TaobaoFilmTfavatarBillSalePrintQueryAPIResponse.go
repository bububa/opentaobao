package filmtfavatar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFilmTfavatarBillSalePrintQueryAPIResponse 获取影院卖品账单-核销账单 API返回值
// taobao.film.tfavatar.bill.sale.print.query
//
// 获取影院卖品账单-核销账单
// 返回值data属于加密字段, 并非大字段.
type TaobaoFilmTfavatarBillSalePrintQueryAPIResponse struct {
	model.CommonResponse
	TaobaoFilmTfavatarBillSalePrintQueryAPIResponseModel
}

// TaobaoFilmTfavatarBillSalePrintQueryAPIResponseModel is 获取影院卖品账单-核销账单 成功返回结果
type TaobaoFilmTfavatarBillSalePrintQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"film_tfavatar_bill_sale_print_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoFilmTfavatarBillSalePrintQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
