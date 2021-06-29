package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
9.1-同步渠道配额 APIRequest
alibaba.tmallgenie.scp.plan.channel.quote.upload

同步渠道配额
*/
type AlibabaTmallgenieScpPlanChannelQuoteUploadRequest struct {
    model.Params

    // 对象
    netDemandRequest   *NetDemandRequest 

}

func NewAlibabaTmallgenieScpPlanChannelQuoteUploadRequest() *AlibabaTmallgenieScpPlanChannelQuoteUploadRequest{
    return &AlibabaTmallgenieScpPlanChannelQuoteUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTmallgenieScpPlanChannelQuoteUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.channel.quote.upload"
}

func (r AlibabaTmallgenieScpPlanChannelQuoteUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTmallgenieScpPlanChannelQuoteUploadRequest) SetNetDemandRequest(netDemandRequest *NetDemandRequest) error {
    r.netDemandRequest = netDemandRequest
    r.Set("net_demand_request", netDemandRequest)
    return nil
}

func (r AlibabaTmallgenieScpPlanChannelQuoteUploadRequest) GetNetDemandRequest() *NetDemandRequest {
    return r.netDemandRequest
}

