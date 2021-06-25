package car

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪订单状态查询接口 APIRequest
taobao.alitrip.car.order.query

提供给直连商家查询在飞猪平台上产生的订单
*/
type TaobaoAlitripCarOrderQueryRequest struct {
    model.Params

    // 飞猪平台订单id
    orderId   string 

}

func NewTaobaoAlitripCarOrderQueryRequest() *TaobaoAlitripCarOrderQueryRequest{
    return &TaobaoAlitripCarOrderQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripCarOrderQueryRequest) GetApiMethodName() string {
    return "taobao.alitrip.car.order.query"
}

func (r TaobaoAlitripCarOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripCarOrderQueryRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r TaobaoAlitripCarOrderQueryRequest) GetOrderId() string {
    return r.orderId
}

