package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenCostCenterModifyAPIRequest 修改成本中心 API请求
// alitrip.btrip.open.cost.center.modify
//
// 修改成本中心
type AlitripBtripOpenCostCenterModifyAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenCostCenterModifyRq
}

// NewAlitripBtripOpenCostCenterModifyRequest 初始化AlitripBtripOpenCostCenterModifyAPIRequest对象
func NewAlitripBtripOpenCostCenterModifyRequest() *AlitripBtripOpenCostCenterModifyAPIRequest {
	return &AlitripBtripOpenCostCenterModifyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripOpenCostCenterModifyAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenCostCenterModifyAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.open.cost.center.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenCostCenterModifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripOpenCostCenterModifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参对象
func (r *AlitripBtripOpenCostCenterModifyAPIRequest) SetRq(_rq *OpenCostCenterModifyRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripOpenCostCenterModifyAPIRequest) GetRq() *OpenCostCenterModifyRq {
	return r._rq
}

var poolAlitripBtripOpenCostCenterModifyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripOpenCostCenterModifyRequest()
	},
}

// GetAlitripBtripOpenCostCenterModifyRequest 从 sync.Pool 获取 AlitripBtripOpenCostCenterModifyAPIRequest
func GetAlitripBtripOpenCostCenterModifyAPIRequest() *AlitripBtripOpenCostCenterModifyAPIRequest {
	return poolAlitripBtripOpenCostCenterModifyAPIRequest.Get().(*AlitripBtripOpenCostCenterModifyAPIRequest)
}

// ReleaseAlitripBtripOpenCostCenterModifyAPIRequest 将 AlitripBtripOpenCostCenterModifyAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripOpenCostCenterModifyAPIRequest(v *AlitripBtripOpenCostCenterModifyAPIRequest) {
	v.Reset()
	poolAlitripBtripOpenCostCenterModifyAPIRequest.Put(v)
}
