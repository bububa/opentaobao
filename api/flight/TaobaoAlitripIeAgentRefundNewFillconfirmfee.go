package flight

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/flight"
)

/* 
新模型-回填申请单费用 
taobao.alitrip.ie.agent.refund.new.fillconfirmfee

1. 回填退票费用
*/
func TaobaoAlitripIeAgentRefundNewFillconfirmfee(clt *core.SDKClient, req *flight.TaobaoAlitripIeAgentRefundNewFillconfirmfeeRequest, session string) (*flight.TaobaoAlitripIeAgentRefundNewFillconfirmfeeResponse, error) {
    var resp flight.TaobaoAlitripIeAgentRefundNewFillconfirmfeeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
