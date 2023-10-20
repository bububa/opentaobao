package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripopencostcenternewAPIRequest 新增成本中心 API请求
// alitrip.btrip.open.cost.center.new
//
// 新增成本中心
type AlitripbtripopencostcenternewAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenCostCenterSaveRq
}

// NewAlitripbtripopencostcenternewRequest 初始化AlitripbtripopencostcenternewAPIRequest对象
func NewAlitripbtripopencostcenternewRequest() *AlitripbtripopencostcenternewAPIRequest {
	return &AlitripbtripopencostcenternewAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripopencostcenternewAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.open.cost.center.new"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripopencostcenternewAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripopencostcenternewAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参对象
func (r *AlitripbtripopencostcenternewAPIRequest) SetRq(_rq *OpenCostCenterSaveRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripopencostcenternewAPIRequest) GetRq() *OpenCostCenterSaveRq {
	return r._rq
}
