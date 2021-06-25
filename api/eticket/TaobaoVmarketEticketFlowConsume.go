package eticket

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/eticket"
)

/* 
无交易类凭证核销 
taobao.vmarket.eticket.flow.consume

无交易类凭证核销
*/
func TaobaoVmarketEticketFlowConsume(clt *core.SDKClient, req *eticket.TaobaoVmarketEticketFlowConsumeRequest, session string) (*eticket.TaobaoVmarketEticketFlowConsumeResponse, error) {
    var resp eticket.TaobaoVmarketEticketFlowConsumeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
