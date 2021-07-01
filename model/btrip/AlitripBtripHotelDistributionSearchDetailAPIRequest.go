package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripHotelDistributionSearchDetailAPIRequest
商旅酒店api分销-详情报价接口 API请求
alitrip.btrip.hotel.distribution.search.detail

商旅酒店api分销-详情报价接口 */
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
func (r AlitripBtripHotelDistributionSearchDetailAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamHotelDetailRQ Setter
// 详情报价入参
func (r *AlitripBtripHotelDistributionSearchDetailAPIRequest) SetParamHotelDetailRQ(_paramHotelDetailRQ *HotelDetailRq) error {
	r._paramHotelDetailRQ = _paramHotelDetailRQ
	r.Set("param_hotel_detail_r_q", _paramHotelDetailRQ)
	return nil
}

// Get ParamHotelDetailRQ Getter
func (r AlitripBtripHotelDistributionSearchDetailAPIRequest) GetParamHotelDetailRQ() *HotelDetailRq {
	return r._paramHotelDetailRQ
}
