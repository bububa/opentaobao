package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
定向推广-计划开启/暂停/删除 
alibaba.scbp.target.ad.plan.operation

定向推广-计划开启/暂停/删除
*/
func AlibabaScbpTargetAdPlanOperation(clt *core.SDKClient, req *scbp.AlibabaScbpTargetAdPlanOperationAPIRequest, session string) (*scbp.AlibabaScbpTargetAdPlanOperationAPIResponse, error) {
    var resp scbp.AlibabaScbpTargetAdPlanOperationAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
