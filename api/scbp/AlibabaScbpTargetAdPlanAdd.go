package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
定向推广-新建计划 
alibaba.scbp.target.ad.plan.add

定向推广-新建单条计划
*/
func AlibabaScbpTargetAdPlanAdd(clt *core.SDKClient, req *scbp.AlibabaScbpTargetAdPlanAddRequest, session string) (*scbp.AlibabaScbpTargetAdPlanAddResponse, error) {
    var resp scbp.AlibabaScbpTargetAdPlanAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
