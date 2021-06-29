package filmtfavatar

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/filmtfavatar"
)

/* 
获取影院票务账单-支付订单 
taobao.film.tfavatar.bill.ticket.payment.query

获取影院票务账单-支付订单
*/
func TaobaoFilmTfavatarBillTicketPaymentQuery(clt *core.SDKClient, req *filmtfavatar.TaobaoFilmTfavatarBillTicketPaymentQueryRequest, session string) (*filmtfavatar.TaobaoFilmTfavatarBillTicketPaymentQueryAPIResponse, error) {
    var resp filmtfavatar.TaobaoFilmTfavatarBillTicketPaymentQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
