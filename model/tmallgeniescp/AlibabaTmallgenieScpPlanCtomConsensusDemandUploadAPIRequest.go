package tmallgeniescp

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIRequest) Reset() {
	r._c2MConsensusDemandRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.ctom.consensus.demand.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTmallgenieScpPlanCtomConsensusDemandUploadRequest()
	},
}

// GetAlibabaTmallgenieScpPlanCtomConsensusDemandUploadRequest 从 sync.Pool 获取 AlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIRequest
func GetAlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIRequest() *AlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIRequest {
	return poolAlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIRequest.Get().(*AlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIRequest)
}

// ReleaseAlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIRequest 将 AlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIRequest(v *AlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIRequest) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanCtomConsensusDemandUploadAPIRequest.Put(v)
}
