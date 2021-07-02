package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanConsensusDemandUploadAPIRequest 20-IBP共识需求回传接口 API请求
// alibaba.tmallgenie.scp.plan.consensus.demand.upload
//
// IBP共识需求回传接口
type AlibabaTmallgenieScpPlanConsensusDemandUploadAPIRequest struct {
	model.Params
	// 入参
	_consensusDemandRequest *ConsensusDemandRequest
}

// NewAlibabaTmallgenieScpPlanConsensusDemandUploadRequest 初始化AlibabaTmallgenieScpPlanConsensusDemandUploadAPIRequest对象
func NewAlibabaTmallgenieScpPlanConsensusDemandUploadRequest() *AlibabaTmallgenieScpPlanConsensusDemandUploadAPIRequest {
	return &AlibabaTmallgenieScpPlanConsensusDemandUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanConsensusDemandUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.consensus.demand.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanConsensusDemandUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ConsensusDemandRequest Setter
// 入参
func (r *AlibabaTmallgenieScpPlanConsensusDemandUploadAPIRequest) SetConsensusDemandRequest(_consensusDemandRequest *ConsensusDemandRequest) error {
	r._consensusDemandRequest = _consensusDemandRequest
	r.Set("consensus_demand_request", _consensusDemandRequest)
	return nil
}

// Get ConsensusDemandRequest Getter
func (r AlibabaTmallgenieScpPlanConsensusDemandUploadAPIRequest) GetConsensusDemandRequest() *ConsensusDemandRequest {
	return r._consensusDemandRequest
}
