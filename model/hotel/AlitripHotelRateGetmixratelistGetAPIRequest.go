package hotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriphotelrategetmixratelistgetAPIRequest 酒店评论接口 API请求
// alitrip.hotel.rate.getmixratelist.get
//
// 酒店评论接口
type AlitriphotelrategetmixratelistgetAPIRequest struct {
	model.Params
	// 评论参数
	_paramGetMixRateListParam *GetMixRateListParam
}

// NewAlitriphotelrategetmixratelistgetRequest 初始化AlitriphotelrategetmixratelistgetAPIRequest对象
func NewAlitriphotelrategetmixratelistgetRequest() *AlitriphotelrategetmixratelistgetAPIRequest {
	return &AlitriphotelrategetmixratelistgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriphotelrategetmixratelistgetAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.rate.getmixratelist.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriphotelrategetmixratelistgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriphotelrategetmixratelistgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamGetMixRateListParam is ParamGetMixRateListParam Setter
// 评论参数
func (r *AlitriphotelrategetmixratelistgetAPIRequest) SetParamGetMixRateListParam(_paramGetMixRateListParam *GetMixRateListParam) error {
	r._paramGetMixRateListParam = _paramGetMixRateListParam
	r.Set("param_get_mix_rate_list_param", _paramGetMixRateListParam)
	return nil
}

// GetParamGetMixRateListParam ParamGetMixRateListParam Getter
func (r AlitriphotelrategetmixratelistgetAPIRequest) GetParamGetMixRateListParam() *GetMixRateListParam {
	return r._paramGetMixRateListParam
}
