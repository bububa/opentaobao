package user

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripFlightchangeAddAPIRequest 航变信息录入接口 API请求
// taobao.alitrip.flightchange.add
//
// 代理商调用航变录入接口,
//
//	简要描述:完成航变信息的自动录入后飞猪机票平台会发生的动作是匹配所有该代理人的订单,如果接口参数指定了飞猪机票订单号，发生的动作是匹配该代理人的指定订单；
//
// 找到与航变航班相关的订单给旅客下发航变短信并出发IVR自动外呼旅客。
type TaobaoAlitripFlightchangeAddAPIRequest struct {
	model.Params
	// 录入参数类
	_flightChangeDataDo *FlightChangeDataDo
}

// NewTaobaoAlitripFlightchangeAddRequest 初始化TaobaoAlitripFlightchangeAddAPIRequest对象
func NewTaobaoAlitripFlightchangeAddRequest() *TaobaoAlitripFlightchangeAddAPIRequest {
	return &TaobaoAlitripFlightchangeAddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripFlightchangeAddAPIRequest) Reset() {
	r._flightChangeDataDo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripFlightchangeAddAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.flightchange.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripFlightchangeAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripFlightchangeAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFlightChangeDataDo is FlightChangeDataDo Setter
// 录入参数类
func (r *TaobaoAlitripFlightchangeAddAPIRequest) SetFlightChangeDataDo(_flightChangeDataDo *FlightChangeDataDo) error {
	r._flightChangeDataDo = _flightChangeDataDo
	r.Set("flight_change_data_do", _flightChangeDataDo)
	return nil
}

// GetFlightChangeDataDo FlightChangeDataDo Getter
func (r TaobaoAlitripFlightchangeAddAPIRequest) GetFlightChangeDataDo() *FlightChangeDataDo {
	return r._flightChangeDataDo
}

var poolTaobaoAlitripFlightchangeAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripFlightchangeAddRequest()
	},
}

// GetTaobaoAlitripFlightchangeAddRequest 从 sync.Pool 获取 TaobaoAlitripFlightchangeAddAPIRequest
func GetTaobaoAlitripFlightchangeAddAPIRequest() *TaobaoAlitripFlightchangeAddAPIRequest {
	return poolTaobaoAlitripFlightchangeAddAPIRequest.Get().(*TaobaoAlitripFlightchangeAddAPIRequest)
}

// ReleaseTaobaoAlitripFlightchangeAddAPIRequest 将 TaobaoAlitripFlightchangeAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripFlightchangeAddAPIRequest(v *TaobaoAlitripFlightchangeAddAPIRequest) {
	v.Reset()
	poolTaobaoAlitripFlightchangeAddAPIRequest.Put(v)
}
