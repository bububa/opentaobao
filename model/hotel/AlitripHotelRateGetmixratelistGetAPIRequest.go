package hotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripHotelRateGetmixratelistGetAPIRequest 酒店评论接口 API请求
// alitrip.hotel.rate.getmixratelist.get
//
// 酒店评论接口
type AlitripHotelRateGetmixratelistGetAPIRequest struct {
	model.Params
	// 评论参数
	_paramGetMixRateListParam *GetMixRateListParam
}

// NewAlitripHotelRateGetmixratelistGetRequest 初始化AlitripHotelRateGetmixratelistGetAPIRequest对象
func NewAlitripHotelRateGetmixratelistGetRequest() *AlitripHotelRateGetmixratelistGetAPIRequest {
	return &AlitripHotelRateGetmixratelistGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripHotelRateGetmixratelistGetAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.rate.getmixratelist.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripHotelRateGetmixratelistGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripHotelRateGetmixratelistGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamGetMixRateListParam is ParamGetMixRateListParam Setter
// 评论参数
func (r *AlitripHotelRateGetmixratelistGetAPIRequest) SetParamGetMixRateListParam(_paramGetMixRateListParam *GetMixRateListParam) error {
	r._paramGetMixRateListParam = _paramGetMixRateListParam
	r.Set("param_get_mix_rate_list_param", _paramGetMixRateListParam)
	return nil
}

// GetParamGetMixRateListParam ParamGetMixRateListParam Getter
func (r AlitripHotelRateGetmixratelistGetAPIRequest) GetParamGetMixRateListParam() *GetMixRateListParam {
	return r._paramGetMixRateListParam
}
