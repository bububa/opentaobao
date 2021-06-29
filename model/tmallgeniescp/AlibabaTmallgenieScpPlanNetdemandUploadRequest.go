package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
23-Net Demand净需求回传 APIRequest
alibaba.tmallgenie.scp.plan.netdemand.upload

Net Demand净需求回传
*/
type AlibabaTmallgenieScpPlanNetdemandUploadRequest struct {
    model.Params

    // 对象
    netDemandRequest   *NetDemandRequest 

}

func NewAlibabaTmallgenieScpPlanNetdemandUploadRequest() *AlibabaTmallgenieScpPlanNetdemandUploadRequest{
    return &AlibabaTmallgenieScpPlanNetdemandUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTmallgenieScpPlanNetdemandUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.netdemand.upload"
}

func (r AlibabaTmallgenieScpPlanNetdemandUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTmallgenieScpPlanNetdemandUploadRequest) SetNetDemandRequest(netDemandRequest *NetDemandRequest) error {
    r.netDemandRequest = netDemandRequest
    r.Set("net_demand_request", netDemandRequest)
    return nil
}

func (r AlibabaTmallgenieScpPlanNetdemandUploadRequest) GetNetDemandRequest() *NetDemandRequest {
    return r.netDemandRequest
}

