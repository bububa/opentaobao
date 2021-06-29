package exchange

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/exchange"
)

/* 
卖家发货 
tmall.exchange.consigngoods

卖家发货
*/
func TmallExchangeConsigngoods(clt *core.SDKClient, req *exchange.TmallExchangeConsigngoodsRequest, session string) (*exchange.TmallExchangeConsigngoodsAPIResponse, error) {
    var resp exchange.TmallExchangeConsigngoodsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
