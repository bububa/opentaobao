package bus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/bus"
)

/* 
出票接口 
taobao.bus.ticket.set

提供给汽车票商家出票使用
*/
func TaobaoBusTicketSet(clt *core.SDKClient, req *bus.TaobaoBusTicketSetRequest, session string) (*bus.TaobaoBusTicketSetResponse, error) {
    var resp bus.TaobaoBusTicketSetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
