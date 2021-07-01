package bus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/bus"
)

/* 
汽车票订单查询 
taobao.bus.order.get

商家汽车票订单查询
*/
func TaobaoBusOrderGet(clt *core.SDKClient, req *bus.TaobaoBusOrderGetAPIRequest, session string) (*bus.TaobaoBusOrderGetAPIResponse, error) {
    var resp bus.TaobaoBusOrderGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
