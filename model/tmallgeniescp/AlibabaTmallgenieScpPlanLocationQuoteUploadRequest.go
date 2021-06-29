package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
9.2-同步地点配额 APIRequest
alibaba.tmallgenie.scp.plan.location.quote.upload

同步地点配额
*/
type AlibabaTmallgenieScpPlanLocationQuoteUploadRequest struct {
    model.Params

    // 对象
    netDemandRequest   *NetDemandRequest 

}

func NewAlibabaTmallgenieScpPlanLocationQuoteUploadRequest() *AlibabaTmallgenieScpPlanLocationQuoteUploadRequest{
    return &AlibabaTmallgenieScpPlanLocationQuoteUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTmallgenieScpPlanLocationQuoteUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.location.quote.upload"
}

func (r AlibabaTmallgenieScpPlanLocationQuoteUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTmallgenieScpPlanLocationQuoteUploadRequest) SetNetDemandRequest(netDemandRequest *NetDemandRequest) error {
    r.netDemandRequest = netDemandRequest
    r.Set("net_demand_request", netDemandRequest)
    return nil
}

func (r AlibabaTmallgenieScpPlanLocationQuoteUploadRequest) GetNetDemandRequest() *NetDemandRequest {
    return r.netDemandRequest
}

