package tmallgeniescp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallgeniescp"
)

/* 
二级物料净需求回传(TL+1) 
alibaba.tmallgenie.scp.plan.netdemand.raw.upload

二级物料净需求回传(TL+1)
*/
func AlibabaTmallgenieScpPlanNetdemandRawUpload(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanNetdemandRawUploadAPIRequest, session string) (*tmallgeniescp.AlibabaTmallgenieScpPlanNetdemandRawUploadAPIResponse, error) {
    var resp tmallgeniescp.AlibabaTmallgenieScpPlanNetdemandRawUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
