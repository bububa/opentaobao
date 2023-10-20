package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenCostCenterEntityDeleteAPIRequest 删除成本中心人员信息 API请求
// alitrip.btrip.open.cost.center.entity.delete
//
// 删除成本中心人员信息
type AlitripBtripOpenCostCenterEntityDeleteAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenCostCenterDeleteEntityRq
}

// NewAlitripBtripOpenCostCenterEntityDeleteRequest 初始化AlitripBtripOpenCostCenterEntityDeleteAPIRequest对象
func NewAlitripBtripOpenCostCenterEntityDeleteRequest() *AlitripBtripOpenCostCenterEntityDeleteAPIRequest {
	return &AlitripBtripOpenCostCenterEntityDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripOpenCostCenterEntityDeleteAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenCostCenterEntityDeleteAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.open.cost.center.entity.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenCostCenterEntityDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripOpenCostCenterEntityDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参对象
func (r *AlitripBtripOpenCostCenterEntityDeleteAPIRequest) SetRq(_rq *OpenCostCenterDeleteEntityRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripOpenCostCenterEntityDeleteAPIRequest) GetRq() *OpenCostCenterDeleteEntityRq {
	return r._rq
}

var poolAlitripBtripOpenCostCenterEntityDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripOpenCostCenterEntityDeleteRequest()
	},
}

// GetAlitripBtripOpenCostCenterEntityDeleteRequest 从 sync.Pool 获取 AlitripBtripOpenCostCenterEntityDeleteAPIRequest
func GetAlitripBtripOpenCostCenterEntityDeleteAPIRequest() *AlitripBtripOpenCostCenterEntityDeleteAPIRequest {
	return poolAlitripBtripOpenCostCenterEntityDeleteAPIRequest.Get().(*AlitripBtripOpenCostCenterEntityDeleteAPIRequest)
}

// ReleaseAlitripBtripOpenCostCenterEntityDeleteAPIRequest 将 AlitripBtripOpenCostCenterEntityDeleteAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripOpenCostCenterEntityDeleteAPIRequest(v *AlitripBtripOpenCostCenterEntityDeleteAPIRequest) {
	v.Reset()
	poolAlitripBtripOpenCostCenterEntityDeleteAPIRequest.Put(v)
}
