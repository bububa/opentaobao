package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenCostCenterEntityAddAPIRequest 增加成本中心人员信息 API请求
// alitrip.btrip.open.cost.center.entity.add
//
// 增加成本中心人员信息
type AlitripBtripOpenCostCenterEntityAddAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenCostCenterAddEntityRq
}

// NewAlitripBtripOpenCostCenterEntityAddRequest 初始化AlitripBtripOpenCostCenterEntityAddAPIRequest对象
func NewAlitripBtripOpenCostCenterEntityAddRequest() *AlitripBtripOpenCostCenterEntityAddAPIRequest {
	return &AlitripBtripOpenCostCenterEntityAddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripOpenCostCenterEntityAddAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenCostCenterEntityAddAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.open.cost.center.entity.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenCostCenterEntityAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripOpenCostCenterEntityAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参对象
func (r *AlitripBtripOpenCostCenterEntityAddAPIRequest) SetRq(_rq *OpenCostCenterAddEntityRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripOpenCostCenterEntityAddAPIRequest) GetRq() *OpenCostCenterAddEntityRq {
	return r._rq
}

var poolAlitripBtripOpenCostCenterEntityAddAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripOpenCostCenterEntityAddRequest()
	},
}

// GetAlitripBtripOpenCostCenterEntityAddRequest 从 sync.Pool 获取 AlitripBtripOpenCostCenterEntityAddAPIRequest
func GetAlitripBtripOpenCostCenterEntityAddAPIRequest() *AlitripBtripOpenCostCenterEntityAddAPIRequest {
	return poolAlitripBtripOpenCostCenterEntityAddAPIRequest.Get().(*AlitripBtripOpenCostCenterEntityAddAPIRequest)
}

// ReleaseAlitripBtripOpenCostCenterEntityAddAPIRequest 将 AlitripBtripOpenCostCenterEntityAddAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripOpenCostCenterEntityAddAPIRequest(v *AlitripBtripOpenCostCenterEntityAddAPIRequest) {
	v.Reset()
	poolAlitripBtripOpenCostCenterEntityAddAPIRequest.Put(v)
}
