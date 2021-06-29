package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同步供应商校准后的配额-二级物料 APIRequest
alibaba.tmallgenie.scp.plan.correct.supplier.quote.raw.upload

同步供应商校准后的配额-二级物料
*/
type AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadRequest struct {
    model.Params

    // 对象
    currentQuoteRawRequest   *AbstractRequest 

}

func NewAlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadRequest() *AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadRequest{
    return &AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.correct.supplier.quote.raw.upload"
}

func (r AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadRequest) SetCurrentQuoteRawRequest(currentQuoteRawRequest *AbstractRequest) error {
    r.currentQuoteRawRequest = currentQuoteRawRequest
    r.Set("current_quote_raw_request", currentQuoteRawRequest)
    return nil
}

func (r AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadRequest) GetCurrentQuoteRawRequest() *AbstractRequest {
    return r.currentQuoteRawRequest
}

