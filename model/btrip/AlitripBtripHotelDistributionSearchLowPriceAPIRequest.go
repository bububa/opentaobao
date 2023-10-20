package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtriphoteldistributionsearchlowpriceAPIRequest 商旅酒店api分销-酒店最低价 API请求
// alitrip.btrip.hotel.distribution.search.low.price
//
// 商旅酒店api分销-酒店最低价
type AlitripbtriphoteldistributionsearchlowpriceAPIRequest struct {
	model.Params
	// 列表最低价入参
	_paramHotelSearchListRQ *HotelSearchListRq
}

// NewAlitripbtriphoteldistributionsearchlowpriceRequest 初始化AlitripbtriphoteldistributionsearchlowpriceAPIRequest对象
func NewAlitripbtriphoteldistributionsearchlowpriceRequest() *AlitripbtriphoteldistributionsearchlowpriceAPIRequest {
	return &AlitripbtriphoteldistributionsearchlowpriceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtriphoteldistributionsearchlowpriceAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.hotel.distribution.search.low.price"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtriphoteldistributionsearchlowpriceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtriphoteldistributionsearchlowpriceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamHotelSearchListRQ is ParamHotelSearchListRQ Setter
// 列表最低价入参
func (r *AlitripbtriphoteldistributionsearchlowpriceAPIRequest) SetParamHotelSearchListRQ(_paramHotelSearchListRQ *HotelSearchListRq) error {
	r._paramHotelSearchListRQ = _paramHotelSearchListRQ
	r.Set("param_hotel_search_list_r_q", _paramHotelSearchListRQ)
	return nil
}

// GetParamHotelSearchListRQ ParamHotelSearchListRQ Getter
func (r AlitripbtriphoteldistributionsearchlowpriceAPIRequest) GetParamHotelSearchListRQ() *HotelSearchListRq {
	return r._paramHotelSearchListRQ
}
