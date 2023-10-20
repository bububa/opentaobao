package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripopencostcentertransferAPIRequest 商旅成本中心转换为外部成本中心 API请求
// alitrip.btrip.open.cost.center.transfer
//
// 商旅成本中心转换为外部成本中心
type AlitripbtripopencostcentertransferAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenCostCenterTransferRq
}

// NewAlitripbtripopencostcentertransferRequest 初始化AlitripbtripopencostcentertransferAPIRequest对象
func NewAlitripbtripopencostcentertransferRequest() *AlitripbtripopencostcentertransferAPIRequest {
	return &AlitripbtripopencostcentertransferAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripopencostcentertransferAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.open.cost.center.transfer"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripopencostcentertransferAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripopencostcentertransferAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参对象
func (r *AlitripbtripopencostcentertransferAPIRequest) SetRq(_rq *OpenCostCenterTransferRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripopencostcentertransferAPIRequest) GetRq() *OpenCostCenterTransferRq {
	return r._rq
}
