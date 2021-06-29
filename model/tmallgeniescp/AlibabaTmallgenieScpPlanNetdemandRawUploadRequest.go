package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
二级物料净需求回传(TL+1) APIRequest
alibaba.tmallgenie.scp.plan.netdemand.raw.upload

二级物料净需求回传(TL+1)
*/
type AlibabaTmallgenieScpPlanNetdemandRawUploadRequest struct {
    model.Params

    // 对象
    netDemandRawRequest   *NetDemandRawRequest 

}

func NewAlibabaTmallgenieScpPlanNetdemandRawUploadRequest() *AlibabaTmallgenieScpPlanNetdemandRawUploadRequest{
    return &AlibabaTmallgenieScpPlanNetdemandRawUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTmallgenieScpPlanNetdemandRawUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.netdemand.raw.upload"
}

func (r AlibabaTmallgenieScpPlanNetdemandRawUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTmallgenieScpPlanNetdemandRawUploadRequest) SetNetDemandRawRequest(netDemandRawRequest *NetDemandRawRequest) error {
    r.netDemandRawRequest = netDemandRawRequest
    r.Set("net_demand_raw_request", netDemandRawRequest)
    return nil
}

func (r AlibabaTmallgenieScpPlanNetdemandRawUploadRequest) GetNetDemandRawRequest() *NetDemandRawRequest {
    return r.netDemandRawRequest
}

