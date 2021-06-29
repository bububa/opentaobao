package iotticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
IoT售后服务商确认接单 API请求
cainiao.iot.ticket.sp.vtwo.accept

IoT售后服务商确认接单
*/
type CainiaoIotTicketSpVtwoAcceptRequest struct {
    model.Params
    // 受理接口请求参数
    acceptTicketTopRequest   *AcceptTicketV2TopRequest
}

// 初始化CainiaoIotTicketSpVtwoAcceptRequest对象
func NewCainiaoIotTicketSpVtwoAcceptRequest() *CainiaoIotTicketSpVtwoAcceptRequest{
    return &CainiaoIotTicketSpVtwoAcceptRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoIotTicketSpVtwoAcceptRequest) GetApiMethodName() string {
    return "cainiao.iot.ticket.sp.vtwo.accept"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoIotTicketSpVtwoAcceptRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AcceptTicketTopRequest Setter
// 受理接口请求参数
func (r *CainiaoIotTicketSpVtwoAcceptRequest) SetAcceptTicketTopRequest(acceptTicketTopRequest *AcceptTicketV2TopRequest) error {
    r.acceptTicketTopRequest = acceptTicketTopRequest
    r.Set("accept_ticket_top_request", acceptTicketTopRequest)
    return nil
}

// AcceptTicketTopRequest Getter
func (r CainiaoIotTicketSpVtwoAcceptRequest) GetAcceptTicketTopRequest() *AcceptTicketV2TopRequest {
    return r.acceptTicketTopRequest
}
