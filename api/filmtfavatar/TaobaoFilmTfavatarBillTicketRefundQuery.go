package filmtfavatar

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/filmtfavatar"
)

/* 
获取影院票务账单-退款账单 
taobao.film.tfavatar.bill.ticket.refund.query

获取影院票务账单-支付订单
data字段为加密字段, 不可分拆.
*/
func TaobaoFilmTfavatarBillTicketRefundQuery(clt *core.SDKClient, req *filmtfavatar.TaobaoFilmTfavatarBillTicketRefundQueryRequest, session string) (*filmtfavatar.TaobaoFilmTfavatarBillTicketRefundQueryAPIResponse, error) {
    var resp filmtfavatar.TaobaoFilmTfavatarBillTicketRefundQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
