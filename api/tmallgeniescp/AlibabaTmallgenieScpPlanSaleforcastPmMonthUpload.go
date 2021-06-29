package tmallgeniescp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallgeniescp"
)

/* 
24-销售月预测数量（产管）回传-月度 
alibaba.tmallgenie.scp.plan.saleforcast.pm.month.upload

销售月预测数量（产管）回传-月度
*/
func AlibabaTmallgenieScpPlanSaleforcastPmMonthUpload(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadRequest, session string) (*tmallgeniescp.AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIResponse, error) {
    var resp tmallgeniescp.AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
