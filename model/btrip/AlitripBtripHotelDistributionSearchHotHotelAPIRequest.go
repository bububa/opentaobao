package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtriphoteldistributionsearchhothotelAPIRequest 商旅酒店api分销-热点酒店 API请求
// alitrip.btrip.hotel.distribution.search.hot.hotel
//
// 商旅酒店api分销-热点酒店
type AlitripbtriphoteldistributionsearchhothotelAPIRequest struct {
	model.Params
	// 列表入参
	_hotHotelSearchListRQ *HotHotelSearchListRq
}

// NewAlitripbtriphoteldistributionsearchhothotelRequest 初始化AlitripbtriphoteldistributionsearchhothotelAPIRequest对象
func NewAlitripbtriphoteldistributionsearchhothotelRequest() *AlitripbtriphoteldistributionsearchhothotelAPIRequest {
	return &AlitripbtriphoteldistributionsearchhothotelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtriphoteldistributionsearchhothotelAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.hotel.distribution.search.hot.hotel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtriphoteldistributionsearchhothotelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtriphoteldistributionsearchhothotelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHotHotelSearchListRQ is HotHotelSearchListRQ Setter
// 列表入参
func (r *AlitripbtriphoteldistributionsearchhothotelAPIRequest) SetHotHotelSearchListRQ(_hotHotelSearchListRQ *HotHotelSearchListRq) error {
	r._hotHotelSearchListRQ = _hotHotelSearchListRQ
	r.Set("hot_hotel_search_list_r_q", _hotHotelSearchListRQ)
	return nil
}

// GetHotHotelSearchListRQ HotHotelSearchListRQ Getter
func (r AlitripbtriphoteldistributionsearchhothotelAPIRequest) GetHotHotelSearchListRQ() *HotHotelSearchListRq {
	return r._hotHotelSearchListRQ
}
