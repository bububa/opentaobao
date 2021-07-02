package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenCostCenterEntityAddAPIRequest 增加成本中心人员信息 API请求
// alitrip.btrip.open.cost.center.entity.add
//
// 增加成本中心人员信息
type AlitripBtripOpenCostCenterEntityAddAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenCostCenterAddEntityRq
}

// NewAlitripBtripOpenCostCenterEntityAddRequest 初始化AlitripBtripOpenCostCenterEntityAddAPIRequest对象
func NewAlitripBtripOpenCostCenterEntityAddRequest() *AlitripBtripOpenCostCenterEntityAddAPIRequest {
	return &AlitripBtripOpenCostCenterEntityAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenCostCenterEntityAddAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.open.cost.center.entity.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenCostCenterEntityAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Rq Setter
// 入参对象
func (r *AlitripBtripOpenCostCenterEntityAddAPIRequest) SetRq(_rq *OpenCostCenterAddEntityRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// Get Rq Getter
func (r AlitripBtripOpenCostCenterEntityAddAPIRequest) GetRq() *OpenCostCenterAddEntityRq {
	return r._rq
}
