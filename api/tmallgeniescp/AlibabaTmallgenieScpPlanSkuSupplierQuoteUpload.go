package tmallgeniescp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallgeniescp"
)

/* 
标准供应商配额同步 
alibaba.tmallgenie.scp.plan.sku.supplier.quote.upload

标准供应商配额同步
*/
func AlibabaTmallgenieScpPlanSkuSupplierQuoteUpload(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadRequest, session string) (*tmallgeniescp.AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIResponse, error) {
    var resp tmallgeniescp.AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
