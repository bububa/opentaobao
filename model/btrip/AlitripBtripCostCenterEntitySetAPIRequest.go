package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripCostCenterEntitySetAPIRequest
设置外部成本中心人员信息 API请求
alitrip.btrip.cost.center.entity.set

设置外部成本中心人员信息 */
type AlitripBtripCostCenterEntitySetAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenCostCenterSetEntityRq
}

// NewAlitripBtripCostCenterEntitySetRequest 初始化AlitripBtripCostCenterEntitySetAPIRequest对象
func NewAlitripBtripCostCenterEntitySetRequest() *AlitripBtripCostCenterEntitySetAPIRequest {
	return &AlitripBtripCostCenterEntitySetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCostCenterEntitySetAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.cost.center.entity.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCostCenterEntitySetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Rq Setter
// 请求对象
func (r *AlitripBtripCostCenterEntitySetAPIRequest) SetRq(_rq *OpenCostCenterSetEntityRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// Get Rq Getter
func (r AlitripBtripCostCenterEntitySetAPIRequest) GetRq() *OpenCostCenterSetEntityRq {
	return r._rq
}
