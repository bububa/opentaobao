package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripCostCenterTransferAPIRequest
商旅成本中心转换为外部成本中心 API请求
alitrip.btrip.cost.center.transfer

商旅成本中心转换为外部成本中心 */
type AlitripBtripCostCenterTransferAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenCostCenterTransferRq
}

// NewAlitripBtripCostCenterTransferRequest 初始化AlitripBtripCostCenterTransferAPIRequest对象
func NewAlitripBtripCostCenterTransferRequest() *AlitripBtripCostCenterTransferAPIRequest {
	return &AlitripBtripCostCenterTransferAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCostCenterTransferAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.cost.center.transfer"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCostCenterTransferAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Rq Setter
// 请求对象
func (r *AlitripBtripCostCenterTransferAPIRequest) SetRq(_rq *OpenCostCenterTransferRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// Get Rq Getter
func (r AlitripBtripCostCenterTransferAPIRequest) GetRq() *OpenCostCenterTransferRq {
	return r._rq
}
