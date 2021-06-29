package iotticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Iot售后服务商取消工单 API请求
cainiao.iot.ticket.sp.cancle

IoT售后服务商取消接单
*/
type CainiaoIotTicketSpCancleRequest struct {
    model.Params
    // 请求参数
    param   *AcceptTicketTopRequest
}

// 初始化CainiaoIotTicketSpCancleRequest对象
func NewCainiaoIotTicketSpCancleRequest() *CainiaoIotTicketSpCancleRequest{
    return &CainiaoIotTicketSpCancleRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoIotTicketSpCancleRequest) GetApiMethodName() string {
    return "cainiao.iot.ticket.sp.cancle"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoIotTicketSpCancleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 请求参数
func (r *CainiaoIotTicketSpCancleRequest) SetParam(param *AcceptTicketTopRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r CainiaoIotTicketSpCancleRequest) GetParam() *AcceptTicketTopRequest {
    return r.param
}
