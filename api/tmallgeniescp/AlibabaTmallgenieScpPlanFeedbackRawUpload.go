package tmallgeniescp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallgeniescp"
)

/* 
15-供应商反馈（原料）同步接口 
alibaba.tmallgenie.scp.plan.feedback.raw.upload

供应商反馈（原料）同步接口
*/
func AlibabaTmallgenieScpPlanFeedbackRawUpload(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest, session string) (*tmallgeniescp.AlibabaTmallgenieScpPlanFeedbackRawUploadAPIResponse, error) {
    var resp tmallgeniescp.AlibabaTmallgenieScpPlanFeedbackRawUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
