package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtipCostCenterQueryAPIRequest 查询外部成本中心 API请求
// alitrip.btip.cost.center.query
//
// 查询外部成本中心
type AlitripBtipCostCenterQueryAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenCostCenterQueryRq
}

// NewAlitripBtipCostCenterQueryRequest 初始化AlitripBtipCostCenterQueryAPIRequest对象
func NewAlitripBtipCostCenterQueryRequest() *AlitripBtipCostCenterQueryAPIRequest {
	return &AlitripBtipCostCenterQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtipCostCenterQueryAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtipCostCenterQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.btip.cost.center.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtipCostCenterQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtipCostCenterQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripBtipCostCenterQueryAPIRequest) SetRq(_rq *OpenCostCenterQueryRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtipCostCenterQueryAPIRequest) GetRq() *OpenCostCenterQueryRq {
	return r._rq
}

var poolAlitripBtipCostCenterQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtipCostCenterQueryRequest()
	},
}

// GetAlitripBtipCostCenterQueryRequest 从 sync.Pool 获取 AlitripBtipCostCenterQueryAPIRequest
func GetAlitripBtipCostCenterQueryAPIRequest() *AlitripBtipCostCenterQueryAPIRequest {
	return poolAlitripBtipCostCenterQueryAPIRequest.Get().(*AlitripBtipCostCenterQueryAPIRequest)
}

// ReleaseAlitripBtipCostCenterQueryAPIRequest 将 AlitripBtipCostCenterQueryAPIRequest 放入 sync.Pool
func ReleaseAlitripBtipCostCenterQueryAPIRequest(v *AlitripBtipCostCenterQueryAPIRequest) {
	v.Reset()
	poolAlitripBtipCostCenterQueryAPIRequest.Put(v)
}
