package tmallgeniescp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallgeniescp"
)

/* 
18-销售预测数量（产管）回传接口 
alibaba.tmallgenie.scp.plan.saleforcast.pm.upload

销售预测数量（产管）回传接口
*/
func AlibabaTmallgenieScpPlanSaleforcastPmUpload(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanSaleforcastPmUploadRequest, session string) (*tmallgeniescp.AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIResponse, error) {
    var resp tmallgeniescp.AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
