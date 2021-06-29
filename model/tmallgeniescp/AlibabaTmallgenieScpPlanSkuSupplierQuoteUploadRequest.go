package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
标准供应商配额同步 API请求
alibaba.tmallgenie.scp.plan.sku.supplier.quote.upload

标准供应商配额同步
*/
type AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadRequest struct {
    model.Params
    // 对象
    _netDemandRequest   *NetDemandRequest
}

// 初始化AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadRequest对象
func NewAlibabaTmallgenieScpPlanSkuSupplierQuoteUploadRequest() *AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadRequest{
    return &AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.sku.supplier.quote.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NetDemandRequest Setter
// 对象
func (r *AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadRequest) SetNetDemandRequest(_netDemandRequest *NetDemandRequest) error {
    r._netDemandRequest = _netDemandRequest
    r.Set("net_demand_request", _netDemandRequest)
    return nil
}

// NetDemandRequest Getter
func (r AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadRequest) GetNetDemandRequest() *NetDemandRequest {
    return r._netDemandRequest
}
