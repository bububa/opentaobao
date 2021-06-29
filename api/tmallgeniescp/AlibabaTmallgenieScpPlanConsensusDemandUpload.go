package tmallgeniescp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallgeniescp"
)

/* 
20-IBP共识需求回传接口 
alibaba.tmallgenie.scp.plan.consensus.demand.upload

IBP共识需求回传接口
*/
func AlibabaTmallgenieScpPlanConsensusDemandUpload(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanConsensusDemandUploadRequest, session string) (*tmallgeniescp.AlibabaTmallgenieScpPlanConsensusDemandUploadAPIResponse, error) {
    var resp tmallgeniescp.AlibabaTmallgenieScpPlanConsensusDemandUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
