package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatmallgeniescpplanconsensusdemanduploadAPIRequest 20-IBP共识需求回传接口 API请求
// alibaba.tmallgenie.scp.plan.consensus.demand.upload
//
// IBP共识需求回传接口
type AlibabatmallgeniescpplanconsensusdemanduploadAPIRequest struct {
	model.Params
	// 入参
	_consensusDemandRequest *ConsensusDemandRequest
}

// NewAlibabatmallgeniescpplanconsensusdemanduploadRequest 初始化AlibabatmallgeniescpplanconsensusdemanduploadAPIRequest对象
func NewAlibabatmallgeniescpplanconsensusdemanduploadRequest() *AlibabatmallgeniescpplanconsensusdemanduploadAPIRequest {
	return &AlibabatmallgeniescpplanconsensusdemanduploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatmallgeniescpplanconsensusdemanduploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.consensus.demand.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatmallgeniescpplanconsensusdemanduploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatmallgeniescpplanconsensusdemanduploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConsensusDemandRequest is ConsensusDemandRequest Setter
// 入参
func (r *AlibabatmallgeniescpplanconsensusdemanduploadAPIRequest) SetConsensusDemandRequest(_consensusDemandRequest *ConsensusDemandRequest) error {
	r._consensusDemandRequest = _consensusDemandRequest
	r.Set("consensus_demand_request", _consensusDemandRequest)
	return nil
}

// GetConsensusDemandRequest ConsensusDemandRequest Getter
func (r AlibabatmallgeniescpplanconsensusdemanduploadAPIRequest) GetConsensusDemandRequest() *ConsensusDemandRequest {
	return r._consensusDemandRequest
}
