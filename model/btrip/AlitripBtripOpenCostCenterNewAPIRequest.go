package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenCostCenterNewAPIRequest 新增成本中心 API请求
// alitrip.btrip.open.cost.center.new
//
// 新增成本中心
type AlitripBtripOpenCostCenterNewAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenCostCenterSaveRq
}

// NewAlitripBtripOpenCostCenterNewRequest 初始化AlitripBtripOpenCostCenterNewAPIRequest对象
func NewAlitripBtripOpenCostCenterNewRequest() *AlitripBtripOpenCostCenterNewAPIRequest {
	return &AlitripBtripOpenCostCenterNewAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripOpenCostCenterNewAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenCostCenterNewAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.open.cost.center.new"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenCostCenterNewAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripOpenCostCenterNewAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参对象
func (r *AlitripBtripOpenCostCenterNewAPIRequest) SetRq(_rq *OpenCostCenterSaveRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripOpenCostCenterNewAPIRequest) GetRq() *OpenCostCenterSaveRq {
	return r._rq
}

var poolAlitripBtripOpenCostCenterNewAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripOpenCostCenterNewRequest()
	},
}

// GetAlitripBtripOpenCostCenterNewRequest 从 sync.Pool 获取 AlitripBtripOpenCostCenterNewAPIRequest
func GetAlitripBtripOpenCostCenterNewAPIRequest() *AlitripBtripOpenCostCenterNewAPIRequest {
	return poolAlitripBtripOpenCostCenterNewAPIRequest.Get().(*AlitripBtripOpenCostCenterNewAPIRequest)
}

// ReleaseAlitripBtripOpenCostCenterNewAPIRequest 将 AlitripBtripOpenCostCenterNewAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripOpenCostCenterNewAPIRequest(v *AlitripBtripOpenCostCenterNewAPIRequest) {
	v.Reset()
	poolAlitripBtripOpenCostCenterNewAPIRequest.Put(v)
}
