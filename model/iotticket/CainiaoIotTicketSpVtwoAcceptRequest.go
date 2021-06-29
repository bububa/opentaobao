package iotticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
IoT售后服务商确认接单 APIRequest
cainiao.iot.ticket.sp.vtwo.accept

IoT售后服务商确认接单
*/
type CainiaoIotTicketSpVtwoAcceptRequest struct {
    model.Params

    // 受理接口请求参数
    acceptTicketTopRequest   *AcceptTicketV2TopRequest 

}

func NewCainiaoIotTicketSpVtwoAcceptRequest() *CainiaoIotTicketSpVtwoAcceptRequest{
    return &CainiaoIotTicketSpVtwoAcceptRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoIotTicketSpVtwoAcceptRequest) GetApiMethodName() string {
    return "cainiao.iot.ticket.sp.vtwo.accept"
}

func (r CainiaoIotTicketSpVtwoAcceptRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoIotTicketSpVtwoAcceptRequest) SetAcceptTicketTopRequest(acceptTicketTopRequest *AcceptTicketV2TopRequest) error {
    r.acceptTicketTopRequest = acceptTicketTopRequest
    r.Set("accept_ticket_top_request", acceptTicketTopRequest)
    return nil
}

func (r CainiaoIotTicketSpVtwoAcceptRequest) GetAcceptTicketTopRequest() *AcceptTicketV2TopRequest {
    return r.acceptTicketTopRequest
}

