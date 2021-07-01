package hotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripHotelDetailStaticinfoGetAPIRequest
详情页静态信息 API请求
alitrip.hotel.detail.staticinfo.get

酒店静态信息TOP接口 */
type AlitripHotelDetailStaticinfoGetAPIRequest struct {
	model.Params
	// 酒店详情页静态信息参数
	_paramHotelTopStaticDetailsParam *HotelTopStaticDetailsParam
}

// NewAlitripHotelDetailStaticinfoGetRequest 初始化AlitripHotelDetailStaticinfoGetAPIRequest对象
func NewAlitripHotelDetailStaticinfoGetRequest() *AlitripHotelDetailStaticinfoGetAPIRequest {
	return &AlitripHotelDetailStaticinfoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripHotelDetailStaticinfoGetAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.detail.staticinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripHotelDetailStaticinfoGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamHotelTopStaticDetailsParam Setter
// 酒店详情页静态信息参数
func (r *AlitripHotelDetailStaticinfoGetAPIRequest) SetParamHotelTopStaticDetailsParam(_paramHotelTopStaticDetailsParam *HotelTopStaticDetailsParam) error {
	r._paramHotelTopStaticDetailsParam = _paramHotelTopStaticDetailsParam
	r.Set("param_hotel_top_static_details_param", _paramHotelTopStaticDetailsParam)
	return nil
}

// Get ParamHotelTopStaticDetailsParam Getter
func (r AlitripHotelDetailStaticinfoGetAPIRequest) GetParamHotelTopStaticDetailsParam() *HotelTopStaticDetailsParam {
	return r._paramHotelTopStaticDetailsParam
}
