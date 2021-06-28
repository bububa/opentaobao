package bus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/bus"
)

/* 
汽车票下单接口 
taobao.bus.order.set

提供给汽车票商家进行下单
*/
func TaobaoBusOrderSet(clt *core.SDKClient, req *bus.TaobaoBusOrderSetRequest, session string) (*bus.TaobaoBusOrderSetAPIResponse, error) {
    var resp bus.TaobaoBusOrderSetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
