package trade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/trade"
)

/* 
延长交易收货时间 
taobao.trade.receivetime.delay

延长交易收货时间
*/
func TaobaoTradeReceivetimeDelay(clt *core.SDKClient, req *trade.TaobaoTradeReceivetimeDelayRequest, session string) (*trade.TaobaoTradeReceivetimeDelayResponse, error) {
    var resp trade.TaobaoTradeReceivetimeDelayAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
