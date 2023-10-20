package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusAgentCityChangeAPIRequest 城市变更 API请求
// taobao.bus.agent.city.change
//
// 代理商通知城市变更，比如可售变为不可售等
type TaobaoBusAgentCityChangeAPIRequest struct {
	model.Params
	// 城市变更请求对象
	_paramCityChangeRQ *CityChangeRq
}

// NewTaobaoBusAgentCityChangeRequest 初始化TaobaoBusAgentCityChangeAPIRequest对象
func NewTaobaoBusAgentCityChangeRequest() *TaobaoBusAgentCityChangeAPIRequest {
	return &TaobaoBusAgentCityChangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusAgentCityChangeAPIRequest) GetApiMethodName() string {
	return "taobao.bus.agent.city.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusAgentCityChangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBusAgentCityChangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCityChangeRQ is ParamCityChangeRQ Setter
// 城市变更请求对象
func (r *TaobaoBusAgentCityChangeAPIRequest) SetParamCityChangeRQ(_paramCityChangeRQ *CityChangeRq) error {
	r._paramCityChangeRQ = _paramCityChangeRQ
	r.Set("param_city_change_r_q", _paramCityChangeRQ)
	return nil
}

// GetParamCityChangeRQ ParamCityChangeRQ Getter
func (r TaobaoBusAgentCityChangeAPIRequest) GetParamCityChangeRQ() *CityChangeRq {
	return r._paramCityChangeRQ
}
