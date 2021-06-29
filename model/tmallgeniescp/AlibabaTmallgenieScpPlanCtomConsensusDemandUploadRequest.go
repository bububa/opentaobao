package tmallgeniescp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
22-C2M共识需求回传接口 API请求
alibaba.tmallgenie.scp.plan.ctom.consensus.demand.upload

C2M 共识需求回传接口
*/
type AlibabaTmallgenieScpPlanCtomConsensusDemandUploadRequest struct {
    model.Params
    // 对象
    c2MConsensusDemandRequest   *C2MConsensusDemandRequest
}

// 初始化AlibabaTmallgenieScpPlanCtomConsensusDemandUploadRequest对象
func NewAlibabaTmallgenieScpPlanCtomConsensusDemandUploadRequest() *AlibabaTmallgenieScpPlanCtomConsensusDemandUploadRequest{
    return &AlibabaTmallgenieScpPlanCtomConsensusDemandUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanCtomConsensusDemandUploadRequest) GetApiMethodName() string {
    return "alibaba.tmallgenie.scp.plan.ctom.consensus.demand.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanCtomConsensusDemandUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// C2MConsensusDemandRequest Setter
// 对象
func (r *AlibabaTmallgenieScpPlanCtomConsensusDemandUploadRequest) SetC2MConsensusDemandRequest(c2MConsensusDemandRequest *C2MConsensusDemandRequest) error {
    r.c2MConsensusDemandRequest = c2MConsensusDemandRequest
    r.Set("c2_m_consensus_demand_request", c2MConsensusDemandRequest)
    return nil
}

// C2MConsensusDemandRequest Getter
func (r AlibabaTmallgenieScpPlanCtomConsensusDemandUploadRequest) GetC2MConsensusDemandRequest() *C2MConsensusDemandRequest {
    return r.c2MConsensusDemandRequest
}
