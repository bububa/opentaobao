package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
定向推广-新增或删除屏蔽词 
alibaba.scbp.target.ad.plan.forbidden.word.modify

定向推广-新增或删除屏蔽词
*/
func AlibabaScbpTargetAdPlanForbiddenWordModify(clt *core.SDKClient, req *scbp.AlibabaScbpTargetAdPlanForbiddenWordModifyRequest, session string) (*scbp.AlibabaScbpTargetAdPlanForbiddenWordModifyAPIResponse, error) {
    var resp scbp.AlibabaScbpTargetAdPlanForbiddenWordModifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
