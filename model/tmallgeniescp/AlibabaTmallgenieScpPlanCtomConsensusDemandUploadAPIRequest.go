package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIRequest 22-C2M共识需求回传接口 API请求
// alibaba.tmallgenie.scp.plan.ctom.consensus.demand.upload
//
// C2M 共识需求回传接口
type AlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIRequest struct {
	model.Params
	// 对象
	_c2MConsensusDemandRequest *C2MConsensusDemandRequest
}

// NewAlibabaTmallgenieScpPlanCtomConsensusDemandUploadRequest 初始化AlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIRequest对象
func NewAlibabaTmallgenieScpPlanCtomConsensusDemandUploadRequest() *AlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIRequest {
	return &AlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.ctom.consensus.demand.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetC2MConsensusDemandRequest is C2MConsensusDemandRequest Setter
// 对象
func (r *AlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIRequest) SetC2MConsensusDemandRequest(_c2MConsensusDemandRequest *C2MConsensusDemandRequest) error {
	r._c2MConsensusDemandRequest = _c2MConsensusDemandRequest
	r.Set("c2_m_consensus_demand_request", _c2MConsensusDemandRequest)
	return nil
}

// GetC2MConsensusDemandRequest C2MConsensusDemandRequest Getter
func (r AlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIRequest) GetC2MConsensusDemandRequest() *C2MConsensusDemandRequest {
	return r._c2MConsensusDemandRequest
}
