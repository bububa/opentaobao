package filmtfavatar

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/filmtfavatar"
)

/* 
获取影院卖品账单--退款账单 
taobao.film.tfavatar.bill.sale.refund.query

获取影院卖品账单--退款账单
*/
func TaobaoFilmTfavatarBillSaleRefundQuery(clt *core.SDKClient, req *filmtfavatar.TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest, session string) (*filmtfavatar.TaobaoFilmTfavatarBillSaleRefundQueryAPIResponse, error) {
    var resp filmtfavatar.TaobaoFilmTfavatarBillSaleRefundQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
