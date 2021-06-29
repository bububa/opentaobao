package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
标准供应商配额同步 APIRequest
alibaba.tmallgenie.scp.plan.sku.supplier.quote.upload

标准供应商配额同步
*/
type AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadRequest struct {
    model.Params

    // 对象
    netDemandRequest   *NetDemandRequest 

}

func NewAlibabaTmallgenieScpPlanSkuSupplierQuoteUploadRequest() *AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadRequest{
    return &AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.sku.supplier.quote.upload"
}

func (r AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadRequest) SetNetDemandRequest(netDemandRequest *NetDemandRequest) error {
    r.netDemandRequest = netDemandRequest
    r.Set("net_demand_request", netDemandRequest)
    return nil
}

func (r AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadRequest) GetNetDemandRequest() *NetDemandRequest {
    return r.netDemandRequest
}

