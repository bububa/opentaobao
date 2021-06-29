package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商校准后的配额同步 API请求
alibaba.tmallgenie.scp.plan.correct.supplier.quote.upload

供应商校准后的配额同步
*/
type AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadRequest struct {
    model.Params
    // 对象
    _netDemandRequest   *NetDemandRequest
}

// 初始化AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadRequest对象
func NewAlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadRequest() *AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadRequest{
    return &AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.correct.supplier.quote.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NetDemandRequest Setter
// 对象
func (r *AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadRequest) SetNetDemandRequest(_netDemandRequest *NetDemandRequest) error {
    r._netDemandRequest = _netDemandRequest
    r.Set("net_demand_request", _netDemandRequest)
    return nil
}

// NetDemandRequest Getter
func (r AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadRequest) GetNetDemandRequest() *NetDemandRequest {
    return r._netDemandRequest
}
