package tmallgeniescp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallgeniescp"
)

/* 
9.2-同步地点配额 
alibaba.tmallgenie.scp.plan.location.quote.upload

同步地点配额
*/
func AlibabaTmallgenieScpPlanLocationQuoteUpload(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanLocationQuoteUploadRequest, session string) (*tmallgeniescp.AlibabaTmallgenieScpPlanLocationQuoteUploadAPIResponse, error) {
    var resp tmallgeniescp.AlibabaTmallgenieScpPlanLocationQuoteUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}