package mos

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mos"
)

/* 
线上小票号获取 
alibaba.mj.oc.online.ticketno.get

线上小票号获取
*/
func AlibabaMjOcOnlineTicketnoGet(clt *core.SDKClient, req *mos.AlibabaMjOcOnlineTicketnoGetAPIRequest, session string) (*mos.AlibabaMjOcOnlineTicketnoGetAPIResponse, error) {
    var resp mos.AlibabaMjOcOnlineTicketnoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
