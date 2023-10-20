package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripopencostcenterentitydeleteAPIRequest 删除成本中心人员信息 API请求
// alitrip.btrip.open.cost.center.entity.delete
//
// 删除成本中心人员信息
type AlitripbtripopencostcenterentitydeleteAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenCostCenterDeleteEntityRq
}

// NewAlitripbtripopencostcenterentitydeleteRequest 初始化AlitripbtripopencostcenterentitydeleteAPIRequest对象
func NewAlitripbtripopencostcenterentitydeleteRequest() *AlitripbtripopencostcenterentitydeleteAPIRequest {
	return &AlitripbtripopencostcenterentitydeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripopencostcenterentitydeleteAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.open.cost.center.entity.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripopencostcenterentitydeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripopencostcenterentitydeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参对象
func (r *AlitripbtripopencostcenterentitydeleteAPIRequest) SetRq(_rq *OpenCostCenterDeleteEntityRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripopencostcenterentitydeleteAPIRequest) GetRq() *OpenCostCenterDeleteEntityRq {
	return r._rq
}
