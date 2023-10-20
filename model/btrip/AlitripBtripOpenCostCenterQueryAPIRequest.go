package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripopencostcenterqueryAPIRequest 查询成本中心 API请求
// alitrip.btrip.open.cost.center.query
//
// 查询成本中心
type AlitripbtripopencostcenterqueryAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenCostCenterQueryRq
}

// NewAlitripbtripopencostcenterqueryRequest 初始化AlitripbtripopencostcenterqueryAPIRequest对象
func NewAlitripbtripopencostcenterqueryRequest() *AlitripbtripopencostcenterqueryAPIRequest {
	return &AlitripbtripopencostcenterqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripopencostcenterqueryAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.open.cost.center.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripopencostcenterqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripopencostcenterqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参对象
func (r *AlitripbtripopencostcenterqueryAPIRequest) SetRq(_rq *OpenCostCenterQueryRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripopencostcenterqueryAPIRequest) GetRq() *OpenCostCenterQueryRq {
	return r._rq
}
