package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
定向推广 按照id操作推广计划的定向溢价标签，包括新增，删除和更新 
alibaba.scbp.target.ad.plan.update.tags

定向推广 按照id操作推广计划的定向溢价标签，包括新增，删除和更新
*/
func AlibabaScbpTargetAdPlanUpdateTags(clt *core.SDKClient, req *scbp.AlibabaScbpTargetAdPlanUpdateTagsRequest, session string) (*scbp.AlibabaScbpTargetAdPlanUpdateTagsAPIResponse, error) {
    var resp scbp.AlibabaScbpTargetAdPlanUpdateTagsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
