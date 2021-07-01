package flight

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/flight"
)

/* 
航变信息录入接口 
taobao.alitrip.flightchange.add

代理商调用航变录入接口,
    简要描述:完成航变信息的自动录入后飞猪机票平台会发生的动作是匹配所有该代理人的订单,如果接口参数指定了飞猪机票订单号，发生的动作是匹配该代理人的指定订单；
找到与航变航班相关的订单给旅客下发航变短信并出发IVR自动外呼旅客。
*/
func TaobaoAlitripFlightchangeAdd(clt *core.SDKClient, req *flight.TaobaoAlitripFlightchangeAddAPIRequest, session string) (*flight.TaobaoAlitripFlightchangeAddAPIResponse, error) {
    var resp flight.TaobaoAlitripFlightchangeAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
