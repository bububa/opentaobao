package car

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/car"
)

/* 
确认接单 
taobao.alitrip.car.order.accept

用来接收服务商确认接单信息
*/
func TaobaoAlitripCarOrderAccept(clt *core.SDKClient, req *car.TaobaoAlitripCarOrderAcceptRequest, session string) (*car.TaobaoAlitripCarOrderAcceptResponse, error) {
    var resp car.TaobaoAlitripCarOrderAcceptAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
