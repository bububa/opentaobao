package btrip

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripHotelDistributionSearchStaticAPIRequest) Reset() {
	r._paramHotelInfoRQ = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripHotelDistributionSearchStaticAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.hotel.distribution.search.static"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripHotelDistributionSearchStaticAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripHotelDistributionSearchStaticAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlitripBtripHotelDistributionSearchStaticAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripHotelDistributionSearchStaticRequest()
	},
}

// GetAlitripBtripHotelDistributionSearchStaticRequest 从 sync.Pool 获取 AlitripBtripHotelDistributionSearchStaticAPIRequest
func GetAlitripBtripHotelDistributionSearchStaticAPIRequest() *AlitripBtripHotelDistributionSearchStaticAPIRequest {
	return poolAlitripBtripHotelDistributionSearchStaticAPIRequest.Get().(*AlitripBtripHotelDistributionSearchStaticAPIRequest)
}

// ReleaseAlitripBtripHotelDistributionSearchStaticAPIRequest 将 AlitripBtripHotelDistributionSearchStaticAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripHotelDistributionSearchStaticAPIRequest(v *AlitripBtripHotelDistributionSearchStaticAPIRequest) {
	v.Reset()
	poolAlitripBtripHotelDistributionSearchStaticAPIRequest.Put(v)
}
