package ieagency

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ieagency"
)

/* 
卖家查询退票申请 
taobao.alitrip.ie.agent.refund.search

卖家查询退票申请
*/
func TaobaoAlitripIeAgentRefundSearch(clt *core.SDKClient, req *ieagency.TaobaoAlitripIeAgentRefundSearchRequest, session string) (*ieagency.TaobaoAlitripIeAgentRefundSearchAPIResponse, error) {
    var resp ieagency.TaobaoAlitripIeAgentRefundSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
