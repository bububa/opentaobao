package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
22-C2M共识需求回传接口 APIRequest
alibaba.tmallgenie.scp.plan.ctom.consensus.demand.upload

C2M 共识需求回传接口
*/
type AlibabaTmallgenieScpPlanCtomConsensusDemandUploadRequest struct {
    model.Params

    // 对象
    c2MConsensusDemandRequest   *C2MConsensusDemandRequest 

}

func NewAlibabaTmallgenieScpPlanCtomConsensusDemandUploadRequest() *AlibabaTmallgenieScpPlanCtomConsensusDemandUploadRequest{
    return &AlibabaTmallgenieScpPlanCtomConsensusDemandUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTmallgenieScpPlanCtomConsensusDemandUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.ctom.consensus.demand.upload"
}

func (r AlibabaTmallgenieScpPlanCtomConsensusDemandUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTmallgenieScpPlanCtomConsensusDemandUploadRequest) SetC2MConsensusDemandRequest(c2MConsensusDemandRequest *C2MConsensusDemandRequest) error {
    r.c2MConsensusDemandRequest = c2MConsensusDemandRequest
    r.Set("c2_m_consensus_demand_request", c2MConsensusDemandRequest)
    return nil
}

func (r AlibabaTmallgenieScpPlanCtomConsensusDemandUploadRequest) GetC2MConsensusDemandRequest() *C2MConsensusDemandRequest {
    return r.c2MConsensusDemandRequest
}

