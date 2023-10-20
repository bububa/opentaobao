package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripSupplychainTrainDetailSearchAPIRequest 【商旅】火车票订单详情查询 API请求
// alitrip.btrip.supplychain.train.detail.search
//
// 【商旅】火车票订单详情查询
type AlitripBtripSupplychainTrainDetailSearchAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenApiSearchDetailRq
}

// NewAlitripBtripSupplychainTrainDetailSearchRequest 初始化AlitripBtripSupplychainTrainDetailSearchAPIRequest对象
func NewAlitripBtripSupplychainTrainDetailSearchRequest() *AlitripBtripSupplychainTrainDetailSearchAPIRequest {
	return &AlitripBtripSupplychainTrainDetailSearchAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripSupplychainTrainDetailSearchAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripSupplychainTrainDetailSearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.supplychain.train.detail.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripSupplychainTrainDetailSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripSupplychainTrainDetailSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripBtripSupplychainTrainDetailSearchAPIRequest) SetRq(_rq *OpenApiSearchDetailRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripSupplychainTrainDetailSearchAPIRequest) GetRq() *OpenApiSearchDetailRq {
	return r._rq
}

var poolAlitripBtripSupplychainTrainDetailSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripSupplychainTrainDetailSearchRequest()
	},
}

// GetAlitripBtripSupplychainTrainDetailSearchRequest 从 sync.Pool 获取 AlitripBtripSupplychainTrainDetailSearchAPIRequest
func GetAlitripBtripSupplychainTrainDetailSearchAPIRequest() *AlitripBtripSupplychainTrainDetailSearchAPIRequest {
	return poolAlitripBtripSupplychainTrainDetailSearchAPIRequest.Get().(*AlitripBtripSupplychainTrainDetailSearchAPIRequest)
}

// ReleaseAlitripBtripSupplychainTrainDetailSearchAPIRequest 将 AlitripBtripSupplychainTrainDetailSearchAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripSupplychainTrainDetailSearchAPIRequest(v *AlitripBtripSupplychainTrainDetailSearchAPIRequest) {
	v.Reset()
	poolAlitripBtripSupplychainTrainDetailSearchAPIRequest.Put(v)
}
