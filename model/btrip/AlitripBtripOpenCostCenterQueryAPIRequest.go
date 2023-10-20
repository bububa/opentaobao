package btrip

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripOpenCostCenterQueryAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenCostCenterQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.open.cost.center.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenCostCenterQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripOpenCostCenterQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlitripBtripOpenCostCenterQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripOpenCostCenterQueryRequest()
	},
}

// GetAlitripBtripOpenCostCenterQueryRequest 从 sync.Pool 获取 AlitripBtripOpenCostCenterQueryAPIRequest
func GetAlitripBtripOpenCostCenterQueryAPIRequest() *AlitripBtripOpenCostCenterQueryAPIRequest {
	return poolAlitripBtripOpenCostCenterQueryAPIRequest.Get().(*AlitripBtripOpenCostCenterQueryAPIRequest)
}

// ReleaseAlitripBtripOpenCostCenterQueryAPIRequest 将 AlitripBtripOpenCostCenterQueryAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripOpenCostCenterQueryAPIRequest(v *AlitripBtripOpenCostCenterQueryAPIRequest) {
	v.Reset()
	poolAlitripBtripOpenCostCenterQueryAPIRequest.Put(v)
}
