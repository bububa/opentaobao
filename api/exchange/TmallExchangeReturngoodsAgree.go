package exchange

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/exchange"
)

/* 
卖家确认收货 
tmall.exchange.returngoods.agree

卖家确认收货
*/
func TmallExchangeReturngoodsAgree(clt *core.SDKClient, req *exchange.TmallExchangeReturngoodsAgreeAPIRequest, session string) (*exchange.TmallExchangeReturngoodsAgreeAPIResponse, error) {
    var resp exchange.TmallExchangeReturngoodsAgreeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
