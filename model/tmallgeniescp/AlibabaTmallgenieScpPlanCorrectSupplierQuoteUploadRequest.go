package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商校准后的配额同步 APIRequest
alibaba.tmallgenie.scp.plan.correct.supplier.quote.upload

供应商校准后的配额同步
*/
type AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadRequest struct {
    model.Params

    // 对象
    netDemandRequest   *NetDemandRequest 

}

func NewAlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadRequest() *AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadRequest{
    return &AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.correct.supplier.quote.upload"
}

func (r AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadRequest) SetNetDemandRequest(netDemandRequest *NetDemandRequest) error {
    r.netDemandRequest = netDemandRequest
    r.Set("net_demand_request", netDemandRequest)
    return nil
}

func (r AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadRequest) GetNetDemandRequest() *NetDemandRequest {
    return r.netDemandRequest
}

