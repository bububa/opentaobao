package exchange

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/exchange"
)

/* 
卖家创建换货留言 
tmall.exchange.message.add

卖家创建换货留言
*/
func TmallExchangeMessageAdd(clt *core.SDKClient, req *exchange.TmallExchangeMessageAddRequest, session string) (*exchange.TmallExchangeMessageAddAPIResponse, error) {
    var resp exchange.TmallExchangeMessageAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
