package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtriphoteldistributionsearchstaticAPIRequest 商旅酒店api分销-酒店静态信息接口 API请求
// alitrip.btrip.hotel.distribution.search.static
//
// 商旅酒店api分销-酒店静态信息接口
type AlitripbtriphoteldistributionsearchstaticAPIRequest struct {
	model.Params
	// 基础信息入参
	_paramHotelInfoRQ *HotelInfoRq
}

// NewAlitripbtriphoteldistributionsearchstaticRequest 初始化AlitripbtriphoteldistributionsearchstaticAPIRequest对象
func NewAlitripbtriphoteldistributionsearchstaticRequest() *AlitripbtriphoteldistributionsearchstaticAPIRequest {
	return &AlitripbtriphoteldistributionsearchstaticAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtriphoteldistributionsearchstaticAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.hotel.distribution.search.static"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtriphoteldistributionsearchstaticAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtriphoteldistributionsearchstaticAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamHotelInfoRQ is ParamHotelInfoRQ Setter
// 基础信息入参
func (r *AlitripbtriphoteldistributionsearchstaticAPIRequest) SetParamHotelInfoRQ(_paramHotelInfoRQ *HotelInfoRq) error {
	r._paramHotelInfoRQ = _paramHotelInfoRQ
	r.Set("param_hotel_info_r_q", _paramHotelInfoRQ)
	return nil
}

// GetParamHotelInfoRQ ParamHotelInfoRQ Getter
func (r AlitripbtriphoteldistributionsearchstaticAPIRequest) GetParamHotelInfoRQ() *HotelInfoRq {
	return r._paramHotelInfoRQ
}
