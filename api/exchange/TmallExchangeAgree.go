package exchange

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/exchange"
)

/* 
卖家同意换货申请 
tmall.exchange.agree

卖家同意换货申请
*/
func TmallExchangeAgree(clt *core.SDKClient, req *exchange.TmallExchangeAgreeAPIRequest, session string) (*exchange.TmallExchangeAgreeAPIResponse, error) {
    var resp exchange.TmallExchangeAgreeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
