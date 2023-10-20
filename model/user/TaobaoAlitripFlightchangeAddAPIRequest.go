package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripflightchangeaddAPIRequest 航变信息录入接口 API请求
// taobao.alitrip.flightchange.add
//
// 代理商调用航变录入接口,
//
//	简要描述:完成航变信息的自动录入后飞猪机票平台会发生的动作是匹配所有该代理人的订单,如果接口参数指定了飞猪机票订单号，发生的动作是匹配该代理人的指定订单；
//
// 找到与航变航班相关的订单给旅客下发航变短信并出发IVR自动外呼旅客。
type TaobaoalitripflightchangeaddAPIRequest struct {
	model.Params
	// 录入参数类
	_flightChangeDataDo *FlightChangeDataDo
}

// NewTaobaoalitripflightchangeaddRequest 初始化TaobaoalitripflightchangeaddAPIRequest对象
func NewTaobaoalitripflightchangeaddRequest() *TaobaoalitripflightchangeaddAPIRequest {
	return &TaobaoalitripflightchangeaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripflightchangeaddAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.flightchange.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripflightchangeaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripflightchangeaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFlightChangeDataDo is FlightChangeDataDo Setter
// 录入参数类
func (r *TaobaoalitripflightchangeaddAPIRequest) SetFlightChangeDataDo(_flightChangeDataDo *FlightChangeDataDo) error {
	r._flightChangeDataDo = _flightChangeDataDo
	r.Set("flight_change_data_do", _flightChangeDataDo)
	return nil
}

// GetFlightChangeDataDo FlightChangeDataDo Getter
func (r TaobaoalitripflightchangeaddAPIRequest) GetFlightChangeDataDo() *FlightChangeDataDo {
	return r._flightChangeDataDo
}
