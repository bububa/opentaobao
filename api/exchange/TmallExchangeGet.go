package exchange

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/exchange"
)

/* 
获取单笔换货详情 
tmall.exchange.get

获取单笔换货详情
*/
func TmallExchangeGet(clt *core.SDKClient, req *exchange.TmallExchangeGetRequest, session string) (*exchange.TmallExchangeGetAPIResponse, error) {
    var resp exchange.TmallExchangeGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
