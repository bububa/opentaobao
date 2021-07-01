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
type AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest struct {
    model.Params
    // 对象
    _netDemandRequest   *NetDemandRequest
}

// 初始化AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest对象
func NewAlibabaTmallgenieScpPlanSkuSupplierQuoteUploadRequest() *AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest{
    return &AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.sku.supplier.quote.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NetDemandRequest Setter
// 对象
func (r *AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest) SetNetDemandRequest(_netDemandRequest *NetDemandRequest) error {
    r._netDemandRequest = _netDemandRequest
    r.Set("net_demand_request", _netDemandRequest)
    return nil
}

// NetDemandRequest Getter
func (r AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest) GetNetDemandRequest() *NetDemandRequest {
    return r._netDemandRequest
}
