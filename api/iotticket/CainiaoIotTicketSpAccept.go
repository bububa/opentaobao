package iotticket

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iotticket"
)

/* 
IoT售后服务商确认接单 
cainiao.iot.ticket.sp.accept

IoT售后服务商确认接单
*/
func CainiaoIotTicketSpAccept(clt *core.SDKClient, req *iotticket.CainiaoIotTicketSpAcceptRequest, session string) (*iotticket.CainiaoIotTicketSpAcceptAPIResponse, error) {
    var resp iotticket.CainiaoIotTicketSpAcceptAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
