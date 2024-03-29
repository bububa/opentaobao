package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripHotelDistributionSearchLowPriceAPIRequest 商旅酒店api分销-酒店最低价 API请求
// alitrip.btrip.hotel.distribution.search.low.price
//
// 商旅酒店api分销-酒店最低价
type AlitripBtripHotelDistributionSearchLowPriceAPIRequest struct {
	model.Params
	// 列表最低价入参
	_paramHotelSearchListRQ *HotelSearchListRq
}

// NewAlitripBtripHotelDistributionSearchLowPriceRequest 初始化AlitripBtripHotelDistributionSearchLowPriceAPIRequest对象
func NewAlitripBtripHotelDistributionSearchLowPriceRequest() *AlitripBtripHotelDistributionSearchLowPriceAPIRequest {
	return &AlitripBtripHotelDistributionSearchLowPriceAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripHotelDistributionSearchLowPriceAPIRequest) Reset() {
	r._paramHotelSearchListRQ = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripHotelDistributionSearchLowPriceAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.hotel.distribution.search.low.price"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripHotelDistributionSearchLowPriceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripHotelDistributionSearchLowPriceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamHotelSearchListRQ is ParamHotelSearchListRQ Setter
// 列表最低价入参
func (r *AlitripBtripHotelDistributionSearchLowPriceAPIRequest) SetParamHotelSearchListRQ(_paramHotelSearchListRQ *HotelSearchListRq) error {
	r._paramHotelSearchListRQ = _paramHotelSearchListRQ
	r.Set("param_hotel_search_list_r_q", _paramHotelSearchListRQ)
	return nil
}

// GetParamHotelSearchListRQ ParamHotelSearchListRQ Getter
func (r AlitripBtripHotelDistributionSearchLowPriceAPIRequest) GetParamHotelSearchListRQ() *HotelSearchListRq {
	return r._paramHotelSearchListRQ
}

var poolAlitripBtripHotelDistributionSearchLowPriceAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripHotelDistributionSearchLowPriceRequest()
	},
}

// GetAlitripBtripHotelDistributionSearchLowPriceRequest 从 sync.Pool 获取 AlitripBtripHotelDistributionSearchLowPriceAPIRequest
func GetAlitripBtripHotelDistributionSearchLowPriceAPIRequest() *AlitripBtripHotelDistributionSearchLowPriceAPIRequest {
	return poolAlitripBtripHotelDistributionSearchLowPriceAPIRequest.Get().(*AlitripBtripHotelDistributionSearchLowPriceAPIRequest)
}

// ReleaseAlitripBtripHotelDistributionSearchLowPriceAPIRequest 将 AlitripBtripHotelDistributionSearchLowPriceAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripHotelDistributionSearchLowPriceAPIRequest(v *AlitripBtripHotelDistributionSearchLowPriceAPIRequest) {
	v.Reset()
	poolAlitripBtripHotelDistributionSearchLowPriceAPIRequest.Put(v)
}
