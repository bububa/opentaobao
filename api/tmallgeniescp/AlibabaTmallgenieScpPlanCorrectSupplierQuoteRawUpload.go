package tmallgeniescp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallgeniescp"
)

/* 
同步供应商校准后的配额-二级物料 
alibaba.tmallgenie.scp.plan.correct.supplier.quote.raw.upload

同步供应商校准后的配额-二级物料
*/
func AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUpload(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest, session string) (*tmallgeniescp.AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIResponse, error) {
    var resp tmallgeniescp.AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
