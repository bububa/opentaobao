package car

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/car"
)

/* 
服务完成API 
taobao.alitrip.car.order.complete

用来接收服务商订单流程完成信息
*/
func TaobaoAlitripCarOrderComplete(clt *core.SDKClient, req *car.TaobaoAlitripCarOrderCompleteRequest, session string) (*car.TaobaoAlitripCarOrderCompleteResponse, error) {
    var resp car.TaobaoAlitripCarOrderCompleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
