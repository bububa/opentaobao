package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripopencostcenterentitysetAPIRequest 设置成本中心人员信息 API请求
// alitrip.btrip.open.cost.center.entity.set
//
// 设置成本中心人员信息
type AlitripbtripopencostcenterentitysetAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenCostCenterSetEntityRq
}

// NewAlitripbtripopencostcenterentitysetRequest 初始化AlitripbtripopencostcenterentitysetAPIRequest对象
func NewAlitripbtripopencostcenterentitysetRequest() *AlitripbtripopencostcenterentitysetAPIRequest {
	return &AlitripbtripopencostcenterentitysetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripopencostcenterentitysetAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.open.cost.center.entity.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripopencostcenterentitysetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripopencostcenterentitysetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参对象
func (r *AlitripbtripopencostcenterentitysetAPIRequest) SetRq(_rq *OpenCostCenterSetEntityRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripopencostcenterentitysetAPIRequest) GetRq() *OpenCostCenterSetEntityRq {
	return r._rq
}
