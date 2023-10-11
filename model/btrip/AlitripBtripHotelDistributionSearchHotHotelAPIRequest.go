package btrip

import (
	"net/url"

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
		Params: model.NewParams(),
	}
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
