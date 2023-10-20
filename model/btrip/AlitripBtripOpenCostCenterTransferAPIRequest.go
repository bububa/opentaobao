package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenCostCenterTransferAPIRequest 商旅成本中心转换为外部成本中心 API请求
// alitrip.btrip.open.cost.center.transfer
//
// 商旅成本中心转换为外部成本中心
type AlitripBtripOpenCostCenterTransferAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenCostCenterTransferRq
}

// NewAlitripBtripOpenCostCenterTransferRequest 初始化AlitripBtripOpenCostCenterTransferAPIRequest对象
func NewAlitripBtripOpenCostCenterTransferRequest() *AlitripBtripOpenCostCenterTransferAPIRequest {
	return &AlitripBtripOpenCostCenterTransferAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripOpenCostCenterTransferAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenCostCenterTransferAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.open.cost.center.transfer"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenCostCenterTransferAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripOpenCostCenterTransferAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参对象
func (r *AlitripBtripOpenCostCenterTransferAPIRequest) SetRq(_rq *OpenCostCenterTransferRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripOpenCostCenterTransferAPIRequest) GetRq() *OpenCostCenterTransferRq {
	return r._rq
}

var poolAlitripBtripOpenCostCenterTransferAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripOpenCostCenterTransferRequest()
	},
}

// GetAlitripBtripOpenCostCenterTransferRequest 从 sync.Pool 获取 AlitripBtripOpenCostCenterTransferAPIRequest
func GetAlitripBtripOpenCostCenterTransferAPIRequest() *AlitripBtripOpenCostCenterTransferAPIRequest {
	return poolAlitripBtripOpenCostCenterTransferAPIRequest.Get().(*AlitripBtripOpenCostCenterTransferAPIRequest)
}

// ReleaseAlitripBtripOpenCostCenterTransferAPIRequest 将 AlitripBtripOpenCostCenterTransferAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripOpenCostCenterTransferAPIRequest(v *AlitripBtripOpenCostCenterTransferAPIRequest) {
	v.Reset()
	poolAlitripBtripOpenCostCenterTransferAPIRequest.Put(v)
}
