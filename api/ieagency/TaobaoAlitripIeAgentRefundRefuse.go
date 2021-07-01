package ieagency

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ieagency"
)

/* 
拒绝退票申请 
taobao.alitrip.ie.agent.refund.refuse

卖家拒绝退票退票申请
*/
func TaobaoAlitripIeAgentRefundRefuse(clt *core.SDKClient, req *ieagency.TaobaoAlitripIeAgentRefundRefuseAPIRequest, session string) (*ieagency.TaobaoAlitripIeAgentRefundRefuseAPIResponse, error) {
    var resp ieagency.TaobaoAlitripIeAgentRefundRefuseAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
