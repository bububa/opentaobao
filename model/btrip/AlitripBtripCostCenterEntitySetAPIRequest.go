package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCostCenterEntitySetAPIRequest 设置外部成本中心人员信息 API请求
// alitrip.btrip.cost.center.entity.set
//
// 设置外部成本中心人员信息
type AlitripBtripCostCenterEntitySetAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenCostCenterSetEntityRq
}

// NewAlitripBtripCostCenterEntitySetRequest 初始化AlitripBtripCostCenterEntitySetAPIRequest对象
func NewAlitripBtripCostCenterEntitySetRequest() *AlitripBtripCostCenterEntitySetAPIRequest {
	return &AlitripBtripCostCenterEntitySetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripCostCenterEntitySetAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCostCenterEntitySetAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.cost.center.entity.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCostCenterEntitySetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripCostCenterEntitySetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripBtripCostCenterEntitySetAPIRequest) SetRq(_rq *OpenCostCenterSetEntityRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripCostCenterEntitySetAPIRequest) GetRq() *OpenCostCenterSetEntityRq {
	return r._rq
}

var poolAlitripBtripCostCenterEntitySetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripCostCenterEntitySetRequest()
	},
}

// GetAlitripBtripCostCenterEntitySetRequest 从 sync.Pool 获取 AlitripBtripCostCenterEntitySetAPIRequest
func GetAlitripBtripCostCenterEntitySetAPIRequest() *AlitripBtripCostCenterEntitySetAPIRequest {
	return poolAlitripBtripCostCenterEntitySetAPIRequest.Get().(*AlitripBtripCostCenterEntitySetAPIRequest)
}

// ReleaseAlitripBtripCostCenterEntitySetAPIRequest 将 AlitripBtripCostCenterEntitySetAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripCostCenterEntitySetAPIRequest(v *AlitripBtripCostCenterEntitySetAPIRequest) {
	v.Reset()
	poolAlitripBtripCostCenterEntitySetAPIRequest.Put(v)
}
