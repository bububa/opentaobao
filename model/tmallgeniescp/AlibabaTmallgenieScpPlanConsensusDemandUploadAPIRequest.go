package tmallgeniescp

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTmallgenieScpPlanConsensusDemandUploadAPIRequest) Reset() {
	r._consensusDemandRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanConsensusDemandUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.consensus.demand.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanConsensusDemandUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTmallgenieScpPlanConsensusDemandUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConsensusDemandRequest is ConsensusDemandRequest Setter
// 入参
func (r *AlibabaTmallgenieScpPlanConsensusDemandUploadAPIRequest) SetConsensusDemandRequest(_consensusDemandRequest *ConsensusDemandRequest) error {
	r._consensusDemandRequest = _consensusDemandRequest
	r.Set("consensus_demand_request", _consensusDemandRequest)
	return nil
}

// GetConsensusDemandRequest ConsensusDemandRequest Getter
func (r AlibabaTmallgenieScpPlanConsensusDemandUploadAPIRequest) GetConsensusDemandRequest() *ConsensusDemandRequest {
	return r._consensusDemandRequest
}

var poolAlibabaTmallgenieScpPlanConsensusDemandUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTmallgenieScpPlanConsensusDemandUploadRequest()
	},
}

// GetAlibabaTmallgenieScpPlanConsensusDemandUploadRequest 从 sync.Pool 获取 AlibabaTmallgenieScpPlanConsensusDemandUploadAPIRequest
func GetAlibabaTmallgenieScpPlanConsensusDemandUploadAPIRequest() *AlibabaTmallgenieScpPlanConsensusDemandUploadAPIRequest {
	return poolAlibabaTmallgenieScpPlanConsensusDemandUploadAPIRequest.Get().(*AlibabaTmallgenieScpPlanConsensusDemandUploadAPIRequest)
}

// ReleaseAlibabaTmallgenieScpPlanConsensusDemandUploadAPIRequest 将 AlibabaTmallgenieScpPlanConsensusDemandUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanConsensusDemandUploadAPIRequest(v *AlibabaTmallgenieScpPlanConsensusDemandUploadAPIRequest) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanConsensusDemandUploadAPIRequest.Put(v)
}
