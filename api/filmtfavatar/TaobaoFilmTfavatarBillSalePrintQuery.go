package filmtfavatar

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/filmtfavatar"
)

/* 
获取影院卖品账单-核销账单 
taobao.film.tfavatar.bill.sale.print.query

获取影院卖品账单-核销账单
返回值data属于加密字段, 并非大字段.
*/
func TaobaoFilmTfavatarBillSalePrintQuery(clt *core.SDKClient, req *filmtfavatar.TaobaoFilmTfavatarBillSalePrintQueryRequest, session string) (*filmtfavatar.TaobaoFilmTfavatarBillSalePrintQueryAPIResponse, error) {
    var resp filmtfavatar.TaobaoFilmTfavatarBillSalePrintQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
