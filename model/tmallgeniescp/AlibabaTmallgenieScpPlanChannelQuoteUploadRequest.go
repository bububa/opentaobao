package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
9.1-同步渠道配额 API请求
alibaba.tmallgenie.scp.plan.channel.quote.upload

同步渠道配额
*/
type AlibabaTmallgenieScpPlanChannelQuoteUploadRequest struct {
    model.Params
    // 对象
    _netDemandRequest   *NetDemandRequest
}

// 初始化AlibabaTmallgenieScpPlanChannelQuoteUploadRequest对象
func NewAlibabaTmallgenieScpPlanChannelQuoteUploadRequest() *AlibabaTmallgenieScpPlanChannelQuoteUploadRequest{
    return &AlibabaTmallgenieScpPlanChannelQuoteUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanChannelQuoteUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.channel.quote.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanChannelQuoteUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NetDemandRequest Setter
// 对象
func (r *AlibabaTmallgenieScpPlanChannelQuoteUploadRequest) SetNetDemandRequest(_netDemandRequest *NetDemandRequest) error {
    r._netDemandRequest = _netDemandRequest
    r.Set("net_demand_request", _netDemandRequest)
    return nil
}

// NetDemandRequest Getter
func (r AlibabaTmallgenieScpPlanChannelQuoteUploadRequest) GetNetDemandRequest() *NetDemandRequest {
    return r._netDemandRequest
}
