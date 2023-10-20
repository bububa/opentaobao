package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCostCenterNewAPIRequest 新建外部成本中心 API请求
// alitrip.btrip.cost.center.new
//
// 新建外部成本中心
type AlitripBtripCostCenterNewAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenCostCenterSaveRq
}

// NewAlitripBtripCostCenterNewRequest 初始化AlitripBtripCostCenterNewAPIRequest对象
func NewAlitripBtripCostCenterNewRequest() *AlitripBtripCostCenterNewAPIRequest {
	return &AlitripBtripCostCenterNewAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripCostCenterNewAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCostCenterNewAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.cost.center.new"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCostCenterNewAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripCostCenterNewAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripBtripCostCenterNewAPIRequest) SetRq(_rq *OpenCostCenterSaveRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripCostCenterNewAPIRequest) GetRq() *OpenCostCenterSaveRq {
	return r._rq
}

var poolAlitripBtripCostCenterNewAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripCostCenterNewRequest()
	},
}

// GetAlitripBtripCostCenterNewRequest 从 sync.Pool 获取 AlitripBtripCostCenterNewAPIRequest
func GetAlitripBtripCostCenterNewAPIRequest() *AlitripBtripCostCenterNewAPIRequest {
	return poolAlitripBtripCostCenterNewAPIRequest.Get().(*AlitripBtripCostCenterNewAPIRequest)
}

// ReleaseAlitripBtripCostCenterNewAPIRequest 将 AlitripBtripCostCenterNewAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripCostCenterNewAPIRequest(v *AlitripBtripCostCenterNewAPIRequest) {
	v.Reset()
	poolAlitripBtripCostCenterNewAPIRequest.Put(v)
}
