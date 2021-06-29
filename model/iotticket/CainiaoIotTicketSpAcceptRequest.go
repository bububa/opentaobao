package iotticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
IoT售后服务商确认接单 APIRequest
cainiao.iot.ticket.sp.accept

IoT售后服务商确认接单
*/
type CainiaoIotTicketSpAcceptRequest struct {
    model.Params

    // 请求参数
    param   *AcceptTicketTopRequest 

}

func NewCainiaoIotTicketSpAcceptRequest() *CainiaoIotTicketSpAcceptRequest{
    return &CainiaoIotTicketSpAcceptRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoIotTicketSpAcceptRequest) GetApiMethodName() string {
    return "cainiao.iot.ticket.sp.accept"
}

func (r CainiaoIotTicketSpAcceptRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoIotTicketSpAcceptRequest) SetParam(param *AcceptTicketTopRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r CainiaoIotTicketSpAcceptRequest) GetParam() *AcceptTicketTopRequest {
    return r.param
}

