package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripTrainOrderSearchAPIRequest 获取火车票订单列表 API请求
// alitrip.btrip.train.order.search
//
// 第三方OA厂商获取自己的火车票数据
type AlitripBtripTrainOrderSearchAPIRequest struct {
	model.Params
	// 请求
	_rq *OpenSearchRq
}

// NewAlitripBtripTrainOrderSearchRequest 初始化AlitripBtripTrainOrderSearchAPIRequest对象
func NewAlitripBtripTrainOrderSearchRequest() *AlitripBtripTrainOrderSearchAPIRequest {
	return &AlitripBtripTrainOrderSearchAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripTrainOrderSearchAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripTrainOrderSearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.train.order.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripTrainOrderSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripTrainOrderSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求
func (r *AlitripBtripTrainOrderSearchAPIRequest) SetRq(_rq *OpenSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripTrainOrderSearchAPIRequest) GetRq() *OpenSearchRq {
	return r._rq
}

var poolAlitripBtripTrainOrderSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripTrainOrderSearchRequest()
	},
}

// GetAlitripBtripTrainOrderSearchRequest 从 sync.Pool 获取 AlitripBtripTrainOrderSearchAPIRequest
func GetAlitripBtripTrainOrderSearchAPIRequest() *AlitripBtripTrainOrderSearchAPIRequest {
	return poolAlitripBtripTrainOrderSearchAPIRequest.Get().(*AlitripBtripTrainOrderSearchAPIRequest)
}

// ReleaseAlitripBtripTrainOrderSearchAPIRequest 将 AlitripBtripTrainOrderSearchAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripTrainOrderSearchAPIRequest(v *AlitripBtripTrainOrderSearchAPIRequest) {
	v.Reset()
	poolAlitripBtripTrainOrderSearchAPIRequest.Put(v)
}
