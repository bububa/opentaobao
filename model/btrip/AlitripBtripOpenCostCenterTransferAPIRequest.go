package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripOpenCostCenterTransferAPIRequest
商旅成本中心转换为外部成本中心 API请求
alitrip.btrip.open.cost.center.transfer

商旅成本中心转换为外部成本中心 */
type AlitripBtripOpenCostCenterTransferAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenCostCenterTransferRq
}

// NewAlitripBtripOpenCostCenterTransferRequest 初始化AlitripBtripOpenCostCenterTransferAPIRequest对象
func NewAlitripBtripOpenCostCenterTransferRequest() *AlitripBtripOpenCostCenterTransferAPIRequest {
	return &AlitripBtripOpenCostCenterTransferAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenCostCenterTransferAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.open.cost.center.transfer"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenCostCenterTransferAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Rq Setter
// 入参对象
func (r *AlitripBtripOpenCostCenterTransferAPIRequest) SetRq(_rq *OpenCostCenterTransferRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// Get Rq Getter
func (r AlitripBtripOpenCostCenterTransferAPIRequest) GetRq() *OpenCostCenterTransferRq {
	return r._rq
}
