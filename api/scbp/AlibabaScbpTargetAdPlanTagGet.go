package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
定向推广-获取计划的定向溢价数据 
alibaba.scbp.target.ad.plan.tag.get

定向推广-获取计划的定向溢价数据
*/
func AlibabaScbpTargetAdPlanTagGet(clt *core.SDKClient, req *scbp.AlibabaScbpTargetAdPlanTagGetAPIRequest, session string) (*scbp.AlibabaScbpTargetAdPlanTagGetAPIResponse, error) {
    var resp scbp.AlibabaScbpTargetAdPlanTagGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
