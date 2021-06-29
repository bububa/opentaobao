package ieagency

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ieagency"
)

/* 
查询申请单详情(新版) 
taobao.alitrip.ie.agent.refund.new.getdetail

查询申请单详情
*/
func TaobaoAlitripIeAgentRefundNewGetdetail(clt *core.SDKClient, req *ieagency.TaobaoAlitripIeAgentRefundNewGetdetailRequest, session string) (*ieagency.TaobaoAlitripIeAgentRefundNewGetdetailAPIResponse, error) {
    var resp ieagency.TaobaoAlitripIeAgentRefundNewGetdetailAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
