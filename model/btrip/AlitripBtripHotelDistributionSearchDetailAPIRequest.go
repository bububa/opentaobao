package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripHotelDistributionSearchDetailAPIRequest 商旅酒店api分销-详情报价接口 API请求
// alitrip.btrip.hotel.distribution.search.detail
//
// 商旅酒店api分销-详情报价接口
type AlitripBtripHotelDistributionSearchDetailAPIRequest struct {
	model.Params
	// 详情报价入参
	_paramHotelDetailRQ *HotelDetailRq
}

// NewAlitripBtripHotelDistributionSearchDetailRequest 初始化AlitripBtripHotelDistributionSearchDetailAPIRequest对象
func NewAlitripBtripHotelDistributionSearchDetailRequest() *AlitripBtripHotelDistributionSearchDetailAPIRequest {
	return &AlitripBtripHotelDistributionSearchDetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripHotelDistributionSearchDetailAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.hotel.distribution.search.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripHotelDistributionSearchDetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripHotelDistributionSearchDetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamHotelDetailRQ is ParamHotelDetailRQ Setter
// 详情报价入参
func (r *AlitripBtripHotelDistributionSearchDetailAPIRequest) SetParamHotelDetailRQ(_paramHotelDetailRQ *HotelDetailRq) error {
	r._paramHotelDetailRQ = _paramHotelDetailRQ
	r.Set("param_hotel_detail_r_q", _paramHotelDetailRQ)
	return nil
}

// GetParamHotelDetailRQ ParamHotelDetailRQ Getter
func (r AlitripBtripHotelDistributionSearchDetailAPIRequest) GetParamHotelDetailRQ() *HotelDetailRq {
	return r._paramHotelDetailRQ
}
