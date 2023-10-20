package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatmallgeniescpplanctomconsensusdemanduploadAPIRequest 22-C2M共识需求回传接口 API请求
// alibaba.tmallgenie.scp.plan.ctom.consensus.demand.upload
//
// C2M 共识需求回传接口
type AlibabatmallgeniescpplanctomconsensusdemanduploadAPIRequest struct {
	model.Params
	// 对象
	_c2MConsensusDemandRequest *C2mconsensusDemandRequest
}

// NewAlibabatmallgeniescpplanctomconsensusdemanduploadRequest 初始化AlibabatmallgeniescpplanctomconsensusdemanduploadAPIRequest对象
func NewAlibabatmallgeniescpplanctomconsensusdemanduploadRequest() *AlibabatmallgeniescpplanctomconsensusdemanduploadAPIRequest {
	return &AlibabatmallgeniescpplanctomconsensusdemanduploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatmallgeniescpplanctomconsensusdemanduploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.ctom.consensus.demand.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatmallgeniescpplanctomconsensusdemanduploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatmallgeniescpplanctomconsensusdemanduploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetC2MConsensusDemandRequest is C2MConsensusDemandRequest Setter
// 对象
func (r *AlibabatmallgeniescpplanctomconsensusdemanduploadAPIRequest) SetC2MConsensusDemandRequest(_c2MConsensusDemandRequest *C2mconsensusDemandRequest) error {
	r._c2MConsensusDemandRequest = _c2MConsensusDemandRequest
	r.Set("c2_m_consensus_demand_request", _c2MConsensusDemandRequest)
	return nil
}

// GetC2MConsensusDemandRequest C2MConsensusDemandRequest Getter
func (r AlibabatmallgeniescpplanctomconsensusdemanduploadAPIRequest) GetC2MConsensusDemandRequest() *C2mconsensusDemandRequest {
	return r._c2MConsensusDemandRequest
}
