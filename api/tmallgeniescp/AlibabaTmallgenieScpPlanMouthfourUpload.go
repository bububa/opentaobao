package tmallgeniescp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallgeniescp"
)

/* 
21-M+4PR 回传接口接口 
alibaba.tmallgenie.scp.plan.mouthfour.upload

M+4 PR 回传接口
*/
func AlibabaTmallgenieScpPlanMouthfourUpload(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanMouthfourUploadRequest, session string) (*tmallgeniescp.AlibabaTmallgenieScpPlanMouthfourUploadAPIResponse, error) {
    var resp tmallgeniescp.AlibabaTmallgenieScpPlanMouthfourUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
