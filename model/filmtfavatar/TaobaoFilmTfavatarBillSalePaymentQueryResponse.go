package filmtfavatar

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取影院卖品账单--支付账单 APIResponse
taobao.film.tfavatar.bill.sale.payment.query

获取影院卖品账单--支付账单
*/
type TaobaoFilmTfavatarBillSalePaymentQueryAPIResponse struct {
    model.CommonResponse
    TaobaoFilmTfavatarBillSalePaymentQueryResponse
}

type TaobaoFilmTfavatarBillSalePaymentQueryResponse struct {
    XMLName xml.Name `xml:"film_tfavatar_bill_sale_payment_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *TaobaoFilmTfavatarBillSalePaymentQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
