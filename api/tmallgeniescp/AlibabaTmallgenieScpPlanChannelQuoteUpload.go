package tmallgeniescp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallgeniescp"
)

/* 
9.1-同步渠道配额 
alibaba.tmallgenie.scp.plan.channel.quote.upload

同步渠道配额
*/
func AlibabaTmallgenieScpPlanChannelQuoteUpload(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanChannelQuoteUploadRequest, session string) (*tmallgeniescp.AlibabaTmallgenieScpPlanChannelQuoteUploadAPIResponse, error) {
    var resp tmallgeniescp.AlibabaTmallgenieScpPlanChannelQuoteUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
