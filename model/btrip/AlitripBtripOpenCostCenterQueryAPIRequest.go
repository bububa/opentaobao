package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenCostCenterQueryAPIRequest 查询成本中心 API请求
// alitrip.btrip.open.cost.center.query
//
// 查询成本中心
type AlitripBtripOpenCostCenterQueryAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenCostCenterQueryRq
}

// NewAlitripBtripOpenCostCenterQueryRequest 初始化AlitripBtripOpenCostCenterQueryAPIRequest对象
func NewAlitripBtripOpenCostCenterQueryRequest() *AlitripBtripOpenCostCenterQueryAPIRequest {
	return &AlitripBtripOpenCostCenterQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenCostCenterQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.open.cost.center.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenCostCenterQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRq is Rq Setter
// 入参对象
func (r *AlitripBtripOpenCostCenterQueryAPIRequest) SetRq(_rq *OpenCostCenterQueryRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripOpenCostCenterQueryAPIRequest) GetRq() *OpenCostCenterQueryRq {
	return r._rq
}
