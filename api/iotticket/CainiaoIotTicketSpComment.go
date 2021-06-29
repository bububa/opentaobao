package iotticket

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/iotticket"
)

/* 
IoT售后服务商工单备注 
cainiao.iot.ticket.sp.comment

IoT售后服务商工单备注
*/
func CainiaoIotTicketSpComment(clt *core.SDKClient, req *iotticket.CainiaoIotTicketSpCommentRequest, session string) (*iotticket.CainiaoIotTicketSpCommentAPIResponse, error) {
    var resp iotticket.CainiaoIotTicketSpCommentAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
