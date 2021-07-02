package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenCostCenterNewAPIRequest 新增成本中心 API请求
// alitrip.btrip.open.cost.center.new
//
// 新增成本中心
type AlitripBtripOpenCostCenterNewAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenCostCenterSaveRq
}

// NewAlitripBtripOpenCostCenterNewRequest 初始化AlitripBtripOpenCostCenterNewAPIRequest对象
func NewAlitripBtripOpenCostCenterNewRequest() *AlitripBtripOpenCostCenterNewAPIRequest {
	return &AlitripBtripOpenCostCenterNewAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenCostCenterNewAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.open.cost.center.new"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenCostCenterNewAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRq is Rq Setter
// 入参对象
func (r *AlitripBtripOpenCostCenterNewAPIRequest) SetRq(_rq *OpenCostCenterSaveRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripOpenCostCenterNewAPIRequest) GetRq() *OpenCostCenterSaveRq {
	return r._rq
}
