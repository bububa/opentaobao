package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
航变信息录入接口 API请求
taobao.alitrip.flightchange.add

代理商调用航变录入接口,
    简要描述:完成航变信息的自动录入后飞猪机票平台会发生的动作是匹配所有该代理人的订单,如果接口参数指定了飞猪机票订单号，发生的动作是匹配该代理人的指定订单；
找到与航变航班相关的订单给旅客下发航变短信并出发IVR自动外呼旅客。
*/
type TaobaoAlitripFlightchangeAddAPIRequest struct {
    model.Params
    // 录入参数类
    _flightChangeDataDo   *FlightChangeDataDo
}

// 初始化TaobaoAlitripFlightchangeAddAPIRequest对象
func NewTaobaoAlitripFlightchangeAddRequest() *TaobaoAlitripFlightchangeAddAPIRequest{
    return &TaobaoAlitripFlightchangeAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripFlightchangeAddAPIRequest) GetApiMethodName() string {
    return "taobao.alitrip.flightchange.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripFlightchangeAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FlightChangeDataDo Setter
// 录入参数类
func (r *TaobaoAlitripFlightchangeAddAPIRequest) SetFlightChangeDataDo(_flightChangeDataDo *FlightChangeDataDo) error {
    r._flightChangeDataDo = _flightChangeDataDo
    r.Set("flight_change_data_do", _flightChangeDataDo)
    return nil
}

// FlightChangeDataDo Getter
func (r TaobaoAlitripFlightchangeAddAPIRequest) GetFlightChangeDataDo() *FlightChangeDataDo {
    return r._flightChangeDataDo
}
