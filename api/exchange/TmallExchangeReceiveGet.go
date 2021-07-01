package exchange

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/exchange"
)

/* 
卖家查询换货列表 
tmall.exchange.receive.get

卖家查询换货列表
*/
func TmallExchangeReceiveGet(clt *core.SDKClient, req *exchange.TmallExchangeReceiveGetAPIRequest, session string) (*exchange.TmallExchangeReceiveGetAPIResponse, error) {
    var resp exchange.TmallExchangeReceiveGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
