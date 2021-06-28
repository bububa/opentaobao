package car

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/car"
)

/* 
租车-取消订单 
taobao.alitrip.car.rent.order.cancel

服务商主动取消用户订单或者拒绝取消订单.
*/
func TaobaoAlitripCarRentOrderCancel(clt *core.SDKClient, req *car.TaobaoAlitripCarRentOrderCancelRequest, session string) (*car.TaobaoAlitripCarRentOrderCancelAPIResponse, error) {
    var resp car.TaobaoAlitripCarRentOrderCancelAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
