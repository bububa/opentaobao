package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobuslastplacegetAPIRequest 获取目的地数据 API请求
// taobao.bus.lastplace.get
//
// 传入城市 获取对应的目的地
type TaobaobuslastplacegetAPIRequest struct {
	model.Params
	// 目的地查询参数
	_paramLastPlaceSearchRQ *ParamLastPlaceSearchRq
}

// NewTaobaobuslastplacegetRequest 初始化TaobaobuslastplacegetAPIRequest对象
func NewTaobaobuslastplacegetRequest() *TaobaobuslastplacegetAPIRequest {
	return &TaobaobuslastplacegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobuslastplacegetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.lastplace.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobuslastplacegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobuslastplacegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamLastPlaceSearchRQ is ParamLastPlaceSearchRQ Setter
// 目的地查询参数
func (r *TaobaobuslastplacegetAPIRequest) SetParamLastPlaceSearchRQ(_paramLastPlaceSearchRQ *ParamLastPlaceSearchRq) error {
	r._paramLastPlaceSearchRQ = _paramLastPlaceSearchRQ
	r.Set("param_last_place_search_r_q", _paramLastPlaceSearchRQ)
	return nil
}

// GetParamLastPlaceSearchRQ ParamLastPlaceSearchRQ Getter
func (r TaobaobuslastplacegetAPIRequest) GetParamLastPlaceSearchRQ() *ParamLastPlaceSearchRq {
	return r._paramLastPlaceSearchRQ
}
