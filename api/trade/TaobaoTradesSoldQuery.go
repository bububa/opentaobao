package trade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/trade"
)

/* 
根据收件人信息查询交易单号 
taobao.trades.sold.query

根据收件人信息查询交易单号。
*/
func TaobaoTradesSoldQuery(clt *core.SDKClient, req *trade.TaobaoTradesSoldQueryAPIRequest, session string) (*trade.TaobaoTradesSoldQueryAPIResponse, error) {
    var resp trade.TaobaoTradesSoldQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
