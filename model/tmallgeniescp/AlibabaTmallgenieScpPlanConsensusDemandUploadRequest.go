package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
20-IBP共识需求回传接口 API请求
alibaba.tmallgenie.scp.plan.consensus.demand.upload

IBP共识需求回传接口
*/
type AlibabaTmallgenieScpPlanConsensusDemandUploadRequest struct {
    model.Params
    // 入参
    consensusDemandRequest   *ConsensusDemandRequest
}

// 初始化AlibabaTmallgenieScpPlanConsensusDemandUploadRequest对象
func NewAlibabaTmallgenieScpPlanConsensusDemandUploadRequest() *AlibabaTmallgenieScpPlanConsensusDemandUploadRequest{
    return &AlibabaTmallgenieScpPlanConsensusDemandUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanConsensusDemandUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.consensus.demand.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanConsensusDemandUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ConsensusDemandRequest Setter
// 入参
func (r *AlibabaTmallgenieScpPlanConsensusDemandUploadRequest) SetConsensusDemandRequest(consensusDemandRequest *ConsensusDemandRequest) error {
    r.consensusDemandRequest = consensusDemandRequest
    r.Set("consensus_demand_request", consensusDemandRequest)
    return nil
}

// ConsensusDemandRequest Getter
func (r AlibabaTmallgenieScpPlanConsensusDemandUploadRequest) GetConsensusDemandRequest() *ConsensusDemandRequest {
    return r.consensusDemandRequest
}
