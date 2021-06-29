package filmtfavatar

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取影院票务账单-退款账单 APIResponse
taobao.film.tfavatar.bill.ticket.refund.query

获取影院票务账单-支付订单
data字段为加密字段, 不可分拆.
*/
type TaobaoFilmTfavatarBillTicketRefundQueryAPIResponse struct {
    model.CommonResponse
    TaobaoFilmTfavatarBillTicketRefundQueryResponse
}

type TaobaoFilmTfavatarBillTicketRefundQueryResponse struct {
    XMLName xml.Name `xml:"film_tfavatar_bill_ticket_refund_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *TaobaoFilmTfavatarBillTicketRefundQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
