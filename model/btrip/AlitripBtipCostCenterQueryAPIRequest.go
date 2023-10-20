package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtipcostcenterqueryAPIRequest 查询外部成本中心 API请求
// alitrip.btip.cost.center.query
//
// 查询外部成本中心
type AlitripbtipcostcenterqueryAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenCostCenterQueryRq
}

// NewAlitripbtipcostcenterqueryRequest 初始化AlitripbtipcostcenterqueryAPIRequest对象
func NewAlitripbtipcostcenterqueryRequest() *AlitripbtipcostcenterqueryAPIRequest {
	return &AlitripbtipcostcenterqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtipcostcenterqueryAPIRequest) GetApiMethodName() string {
	return "alitrip.btip.cost.center.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtipcostcenterqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtipcostcenterqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripbtipcostcenterqueryAPIRequest) SetRq(_rq *OpenCostCenterQueryRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtipcostcenterqueryAPIRequest) GetRq() *OpenCostCenterQueryRq {
	return r._rq
}
