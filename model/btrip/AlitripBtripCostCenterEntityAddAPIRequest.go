package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCostCenterEntityAddAPIRequest 增加外部成本中心人员信息 API请求
// alitrip.btrip.cost.center.entity.add
//
// 增加外部成本中心人员信息
type AlitripBtripCostCenterEntityAddAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenCostCenterAddEntityRq
}

// NewAlitripBtripCostCenterEntityAddRequest 初始化AlitripBtripCostCenterEntityAddAPIRequest对象
func NewAlitripBtripCostCenterEntityAddRequest() *AlitripBtripCostCenterEntityAddAPIRequest {
	return &AlitripBtripCostCenterEntityAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCostCenterEntityAddAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.cost.center.entity.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCostCenterEntityAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripCostCenterEntityAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripBtripCostCenterEntityAddAPIRequest) SetRq(_rq *OpenCostCenterAddEntityRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripCostCenterEntityAddAPIRequest) GetRq() *OpenCostCenterAddEntityRq {
	return r._rq
}
