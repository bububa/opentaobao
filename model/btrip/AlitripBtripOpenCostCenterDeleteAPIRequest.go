package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripopencostcenterdeleteAPIRequest 删除成本中心 API请求
// alitrip.btrip.open.cost.center.delete
//
// 删除成本中心
type AlitripbtripopencostcenterdeleteAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenCostCenterDeleteRq
}

// NewAlitripbtripopencostcenterdeleteRequest 初始化AlitripbtripopencostcenterdeleteAPIRequest对象
func NewAlitripbtripopencostcenterdeleteRequest() *AlitripbtripopencostcenterdeleteAPIRequest {
	return &AlitripbtripopencostcenterdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripopencostcenterdeleteAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.open.cost.center.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripopencostcenterdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripopencostcenterdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参对象
func (r *AlitripbtripopencostcenterdeleteAPIRequest) SetRq(_rq *OpenCostCenterDeleteRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripopencostcenterdeleteAPIRequest) GetRq() *OpenCostCenterDeleteRq {
	return r._rq
}
