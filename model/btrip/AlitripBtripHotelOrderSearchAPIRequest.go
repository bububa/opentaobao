package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripHotelOrderSearchAPIRequest 搜索酒店订单列表 API请求
// alitrip.btrip.hotel.order.search
//
// 企业获取商旅酒店订单数据
type AlitripBtripHotelOrderSearchAPIRequest struct {
	model.Params
	// 请求
	_rq *OpenSearchRq
}

// NewAlitripBtripHotelOrderSearchRequest 初始化AlitripBtripHotelOrderSearchAPIRequest对象
func NewAlitripBtripHotelOrderSearchRequest() *AlitripBtripHotelOrderSearchAPIRequest {
	return &AlitripBtripHotelOrderSearchAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripHotelOrderSearchAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripHotelOrderSearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.hotel.order.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripHotelOrderSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripHotelOrderSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求
func (r *AlitripBtripHotelOrderSearchAPIRequest) SetRq(_rq *OpenSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripHotelOrderSearchAPIRequest) GetRq() *OpenSearchRq {
	return r._rq
}

var poolAlitripBtripHotelOrderSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripHotelOrderSearchRequest()
	},
}

// GetAlitripBtripHotelOrderSearchRequest 从 sync.Pool 获取 AlitripBtripHotelOrderSearchAPIRequest
func GetAlitripBtripHotelOrderSearchAPIRequest() *AlitripBtripHotelOrderSearchAPIRequest {
	return poolAlitripBtripHotelOrderSearchAPIRequest.Get().(*AlitripBtripHotelOrderSearchAPIRequest)
}

// ReleaseAlitripBtripHotelOrderSearchAPIRequest 将 AlitripBtripHotelOrderSearchAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripHotelOrderSearchAPIRequest(v *AlitripBtripHotelOrderSearchAPIRequest) {
	v.Reset()
	poolAlitripBtripHotelOrderSearchAPIRequest.Put(v)
}
