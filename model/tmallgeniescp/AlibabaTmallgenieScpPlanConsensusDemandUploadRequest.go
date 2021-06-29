package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
20-IBP共识需求回传接口 APIRequest
alibaba.tmallgenie.scp.plan.consensus.demand.upload

IBP共识需求回传接口
*/
type AlibabaTmallgenieScpPlanConsensusDemandUploadRequest struct {
    model.Params

    // 入参
    consensusDemandRequest   *ConsensusDemandRequest 

}

func NewAlibabaTmallgenieScpPlanConsensusDemandUploadRequest() *AlibabaTmallgenieScpPlanConsensusDemandUploadRequest{
    return &AlibabaTmallgenieScpPlanConsensusDemandUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTmallgenieScpPlanConsensusDemandUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.consensus.demand.upload"
}

func (r AlibabaTmallgenieScpPlanConsensusDemandUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTmallgenieScpPlanConsensusDemandUploadRequest) SetConsensusDemandRequest(consensusDemandRequest *ConsensusDemandRequest) error {
    r.consensusDemandRequest = consensusDemandRequest
    r.Set("consensus_demand_request", consensusDemandRequest)
    return nil
}

func (r AlibabaTmallgenieScpPlanConsensusDemandUploadRequest) GetConsensusDemandRequest() *ConsensusDemandRequest {
    return r.consensusDemandRequest
}

