package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
9.2-同步地点配额 API请求
alibaba.tmallgenie.scp.plan.location.quote.upload

同步地点配额
*/
type AlibabaTmallgenieScpPlanLocationQuoteUploadRequest struct {
    model.Params
    // 对象
    _netDemandRequest   *NetDemandRequest
}

// 初始化AlibabaTmallgenieScpPlanLocationQuoteUploadRequest对象
func NewAlibabaTmallgenieScpPlanLocationQuoteUploadRequest() *AlibabaTmallgenieScpPlanLocationQuoteUploadRequest{
    return &AlibabaTmallgenieScpPlanLocationQuoteUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanLocationQuoteUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.location.quote.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanLocationQuoteUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NetDemandRequest Setter
// 对象
func (r *AlibabaTmallgenieScpPlanLocationQuoteUploadRequest) SetNetDemandRequest(_netDemandRequest *NetDemandRequest) error {
    r._netDemandRequest = _netDemandRequest
    r.Set("net_demand_request", _netDemandRequest)
    return nil
}

// NetDemandRequest Getter
func (r AlibabaTmallgenieScpPlanLocationQuoteUploadRequest) GetNetDemandRequest() *NetDemandRequest {
    return r._netDemandRequest
}
