package iotticket

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iotticket"
)

/* 
Iot售后服务商取消工单 
cainiao.iot.ticket.sp.cancle

IoT售后服务商取消接单
*/
func CainiaoIotTicketSpCancle(clt *core.SDKClient, req *iotticket.CainiaoIotTicketSpCancleRequest, session string) (*iotticket.CainiaoIotTicketSpCancleAPIResponse, error) {
    var resp iotticket.CainiaoIotTicketSpCancleAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}