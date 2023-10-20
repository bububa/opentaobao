package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripSupplychainTrainCityAPIRequest 火车站数据查询 API请求
// alitrip.btrip.supplychain.train.city
//
// 火车站数据查询
type AlitripBtripSupplychainTrainCityAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenSuggestRq
}

// NewAlitripBtripSupplychainTrainCityRequest 初始化AlitripBtripSupplychainTrainCityAPIRequest对象
func NewAlitripBtripSupplychainTrainCityRequest() *AlitripBtripSupplychainTrainCityAPIRequest {
	return &AlitripBtripSupplychainTrainCityAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripSupplychainTrainCityAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripSupplychainTrainCityAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.supplychain.train.city"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripSupplychainTrainCityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripSupplychainTrainCityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参对象
func (r *AlitripBtripSupplychainTrainCityAPIRequest) SetRq(_rq *OpenSuggestRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripSupplychainTrainCityAPIRequest) GetRq() *OpenSuggestRq {
	return r._rq
}

var poolAlitripBtripSupplychainTrainCityAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripSupplychainTrainCityRequest()
	},
}

// GetAlitripBtripSupplychainTrainCityRequest 从 sync.Pool 获取 AlitripBtripSupplychainTrainCityAPIRequest
func GetAlitripBtripSupplychainTrainCityAPIRequest() *AlitripBtripSupplychainTrainCityAPIRequest {
	return poolAlitripBtripSupplychainTrainCityAPIRequest.Get().(*AlitripBtripSupplychainTrainCityAPIRequest)
}

// ReleaseAlitripBtripSupplychainTrainCityAPIRequest 将 AlitripBtripSupplychainTrainCityAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripSupplychainTrainCityAPIRequest(v *AlitripBtripSupplychainTrainCityAPIRequest) {
	v.Reset()
	poolAlitripBtripSupplychainTrainCityAPIRequest.Put(v)
}
