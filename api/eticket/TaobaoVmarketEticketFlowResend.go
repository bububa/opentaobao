package eticket

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/eticket"
)

/* 
业务重新触发发码短信 
taobao.vmarket.eticket.flow.resend

业务重新触发发码短信
*/
func TaobaoVmarketEticketFlowResend(clt *core.SDKClient, req *eticket.TaobaoVmarketEticketFlowResendRequest, session string) (*eticket.TaobaoVmarketEticketFlowResendAPIResponse, error) {
    var resp eticket.TaobaoVmarketEticketFlowResendAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
