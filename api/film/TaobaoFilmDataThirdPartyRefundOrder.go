package film

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/film"
)

/* 
退票接口 
taobao.film.data.third.party.refund.order

淘票票第三方退票接口
*/
func TaobaoFilmDataThirdPartyRefundOrder(clt *core.SDKClient, req *film.TaobaoFilmDataThirdPartyRefundOrderRequest, session string) (*film.TaobaoFilmDataThirdPartyRefundOrderResponse, error) {
    var resp film.TaobaoFilmDataThirdPartyRefundOrderAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
