package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
【商旅】更新审批单状态 
alitrip.btrip.corpop.apply.approve

【商旅】更新审批单状态
*/
func AlitripBtripCorpopApplyApprove(clt *core.SDKClient, req *btrip.AlitripBtripCorpopApplyApproveAPIRequest, session string) (*btrip.AlitripBtripCorpopApplyApproveAPIResponse, error) {
    var resp btrip.AlitripBtripCorpopApplyApproveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
