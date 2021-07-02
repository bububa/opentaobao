package btrip

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripHotelDistributionSearchLowPriceAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.hotel.distribution.search.low.price"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripHotelDistributionSearchLowPriceAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
