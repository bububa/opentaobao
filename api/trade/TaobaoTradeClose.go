package trade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/trade"
)

/* 
卖家关闭一笔交易 
taobao.trade.close

关闭一笔订单，可以是主订单或子订单。当订单从创建到关闭时间小于10s的时候，会报“CLOSE_TRADE_TOO_FAST”错误。
*/
func TaobaoTradeClose(clt *core.SDKClient, req *trade.TaobaoTradeCloseRequest, session string) (*trade.TaobaoTradeCloseResponse, error) {
    var resp trade.TaobaoTradeCloseAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
