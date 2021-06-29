package filmtfavatar

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取影院票务账单-支付订单 APIResponse
taobao.film.tfavatar.bill.ticket.payment.query

获取影院票务账单-支付订单
*/
type TaobaoFilmTfavatarBillTicketPaymentQueryAPIResponse struct {
    model.CommonResponse
    TaobaoFilmTfavatarBillTicketPaymentQueryResponse
}

type TaobaoFilmTfavatarBillTicketPaymentQueryResponse struct {
    XMLName xml.Name `xml:"film_tfavatar_bill_ticket_payment_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *TaobaoFilmTfavatarBillTicketPaymentQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
