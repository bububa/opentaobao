package bus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/bus"
)

/* 
商家top回调退款明细 
taobao.bus.agent.refundticket.confirm

商家通过top回调告知平台退款明细
*/
func TaobaoBusAgentRefundticketConfirm(clt *core.SDKClient, req *bus.TaobaoBusAgentRefundticketConfirmRequest, session string) (*bus.TaobaoBusAgentRefundticketConfirmAPIResponse, error) {
    var resp bus.TaobaoBusAgentRefundticketConfirmAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
