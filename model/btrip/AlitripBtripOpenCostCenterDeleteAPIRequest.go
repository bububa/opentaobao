package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenCostCenterDeleteAPIRequest 删除成本中心 API请求
// alitrip.btrip.open.cost.center.delete
//
// 删除成本中心
type AlitripBtripOpenCostCenterDeleteAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenCostCenterDeleteRq
}

// NewAlitripBtripOpenCostCenterDeleteRequest 初始化AlitripBtripOpenCostCenterDeleteAPIRequest对象
func NewAlitripBtripOpenCostCenterDeleteRequest() *AlitripBtripOpenCostCenterDeleteAPIRequest {
	return &AlitripBtripOpenCostCenterDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripOpenCostCenterDeleteAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenCostCenterDeleteAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.open.cost.center.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenCostCenterDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripOpenCostCenterDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参对象
func (r *AlitripBtripOpenCostCenterDeleteAPIRequest) SetRq(_rq *OpenCostCenterDeleteRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripOpenCostCenterDeleteAPIRequest) GetRq() *OpenCostCenterDeleteRq {
	return r._rq
}

var poolAlitripBtripOpenCostCenterDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripOpenCostCenterDeleteRequest()
	},
}

// GetAlitripBtripOpenCostCenterDeleteRequest 从 sync.Pool 获取 AlitripBtripOpenCostCenterDeleteAPIRequest
func GetAlitripBtripOpenCostCenterDeleteAPIRequest() *AlitripBtripOpenCostCenterDeleteAPIRequest {
	return poolAlitripBtripOpenCostCenterDeleteAPIRequest.Get().(*AlitripBtripOpenCostCenterDeleteAPIRequest)
}

// ReleaseAlitripBtripOpenCostCenterDeleteAPIRequest 将 AlitripBtripOpenCostCenterDeleteAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripOpenCostCenterDeleteAPIRequest(v *AlitripBtripOpenCostCenterDeleteAPIRequest) {
	v.Reset()
	poolAlitripBtripOpenCostCenterDeleteAPIRequest.Put(v)
}
