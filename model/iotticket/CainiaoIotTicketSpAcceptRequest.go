package iotticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
IoT售后服务商确认接单 API请求
cainiao.iot.ticket.sp.accept

IoT售后服务商确认接单
*/
type CainiaoIotTicketSpAcceptRequest struct {
    model.Params
    // 请求参数
    _param   *AcceptTicketTopRequest
}

// 初始化CainiaoIotTicketSpAcceptRequest对象
func NewCainiaoIotTicketSpAcceptRequest() *CainiaoIotTicketSpAcceptRequest{
    return &CainiaoIotTicketSpAcceptRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoIotTicketSpAcceptRequest) GetApiMethodName() string {
    return "cainiao.iot.ticket.sp.accept"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoIotTicketSpAcceptRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 请求参数
func (r *CainiaoIotTicketSpAcceptRequest) SetParam(_param *AcceptTicketTopRequest) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r CainiaoIotTicketSpAcceptRequest) GetParam() *AcceptTicketTopRequest {
    return r._param
}
