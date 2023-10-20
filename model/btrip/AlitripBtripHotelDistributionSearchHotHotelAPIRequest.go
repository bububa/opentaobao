package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripHotelDistributionSearchHotHotelAPIRequest 商旅酒店api分销-热点酒店 API请求
// alitrip.btrip.hotel.distribution.search.hot.hotel
//
// 商旅酒店api分销-热点酒店
type AlitripBtripHotelDistributionSearchHotHotelAPIRequest struct {
	model.Params
	// 列表入参
	_hotHotelSearchListRQ *HotHotelSearchListRq
}

// NewAlitripBtripHotelDistributionSearchHotHotelRequest 初始化AlitripBtripHotelDistributionSearchHotHotelAPIRequest对象
func NewAlitripBtripHotelDistributionSearchHotHotelRequest() *AlitripBtripHotelDistributionSearchHotHotelAPIRequest {
	return &AlitripBtripHotelDistributionSearchHotHotelAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripHotelDistributionSearchHotHotelAPIRequest) Reset() {
	r._hotHotelSearchListRQ = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripHotelDistributionSearchHotHotelAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.hotel.distribution.search.hot.hotel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripHotelDistributionSearchHotHotelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripHotelDistributionSearchHotHotelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHotHotelSearchListRQ is HotHotelSearchListRQ Setter
// 列表入参
func (r *AlitripBtripHotelDistributionSearchHotHotelAPIRequest) SetHotHotelSearchListRQ(_hotHotelSearchListRQ *HotHotelSearchListRq) error {
	r._hotHotelSearchListRQ = _hotHotelSearchListRQ
	r.Set("hot_hotel_search_list_r_q", _hotHotelSearchListRQ)
	return nil
}

// GetHotHotelSearchListRQ HotHotelSearchListRQ Getter
func (r AlitripBtripHotelDistributionSearchHotHotelAPIRequest) GetHotHotelSearchListRQ() *HotHotelSearchListRq {
	return r._hotHotelSearchListRQ
}

var poolAlitripBtripHotelDistributionSearchHotHotelAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripHotelDistributionSearchHotHotelRequest()
	},
}

// GetAlitripBtripHotelDistributionSearchHotHotelRequest 从 sync.Pool 获取 AlitripBtripHotelDistributionSearchHotHotelAPIRequest
func GetAlitripBtripHotelDistributionSearchHotHotelAPIRequest() *AlitripBtripHotelDistributionSearchHotHotelAPIRequest {
	return poolAlitripBtripHotelDistributionSearchHotHotelAPIRequest.Get().(*AlitripBtripHotelDistributionSearchHotHotelAPIRequest)
}

// ReleaseAlitripBtripHotelDistributionSearchHotHotelAPIRequest 将 AlitripBtripHotelDistributionSearchHotHotelAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripHotelDistributionSearchHotHotelAPIRequest(v *AlitripBtripHotelDistributionSearchHotHotelAPIRequest) {
	v.Reset()
	poolAlitripBtripHotelDistributionSearchHotHotelAPIRequest.Put(v)
}
