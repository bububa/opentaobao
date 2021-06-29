package iotticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Iot售后服务商取消工单 APIRequest
cainiao.iot.ticket.sp.cancle

IoT售后服务商取消接单
*/
type CainiaoIotTicketSpCancleRequest struct {
    model.Params

    // 请求参数
    param   *AcceptTicketTopRequest 

}

func NewCainiaoIotTicketSpCancleRequest() *CainiaoIotTicketSpCancleRequest{
    return &CainiaoIotTicketSpCancleRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoIotTicketSpCancleRequest) GetApiMethodName() string {
    return "cainiao.iot.ticket.sp.cancle"
}

func (r CainiaoIotTicketSpCancleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoIotTicketSpCancleRequest) SetParam(param *AcceptTicketTopRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r CainiaoIotTicketSpCancleRequest) GetParam() *AcceptTicketTopRequest {
    return r.param
}

