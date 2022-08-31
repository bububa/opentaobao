package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripHotelDistributionSearchStaticAPIRequest 商旅酒店api分销-酒店静态信息接口 API请求
// alitrip.btrip.hotel.distribution.search.static
//
// 商旅酒店api分销-酒店静态信息接口
type AlitripBtripHotelDistributionSearchStaticAPIRequest struct {
	model.Params
	// 基础信息入参
	_paramHotelInfoRQ *HotelInfoRq
}

// NewAlitripBtripHotelDistributionSearchStaticRequest 初始化AlitripBtripHotelDistributionSearchStaticAPIRequest对象
func NewAlitripBtripHotelDistributionSearchStaticRequest() *AlitripBtripHotelDistributionSearchStaticAPIRequest {
	return &AlitripBtripHotelDistributionSearchStaticAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripHotelDistributionSearchStaticAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.hotel.distribution.search.static"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripHotelDistributionSearchStaticAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamHotelInfoRQ is ParamHotelInfoRQ Setter
// 基础信息入参
func (r *AlitripBtripHotelDistributionSearchStaticAPIRequest) SetParamHotelInfoRQ(_paramHotelInfoRQ *HotelInfoRq) error {
	r._paramHotelInfoRQ = _paramHotelInfoRQ
	r.Set("param_hotel_info_r_q", _paramHotelInfoRQ)
	return nil
}

// GetParamHotelInfoRQ ParamHotelInfoRQ Getter
func (r AlitripBtripHotelDistributionSearchStaticAPIRequest) GetParamHotelInfoRQ() *HotelInfoRq {
	return r._paramHotelInfoRQ
}
