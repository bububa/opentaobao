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
type CainiaoIotTicketSpVtwoAcceptAPIRequest struct {
    model.Params
    // 受理接口请求参数
    _acceptTicketTopRequest   *AcceptTicketV2TopRequest
}

// 初始化CainiaoIotTicketSpVtwoAcceptAPIRequest对象
func NewCainiaoIotTicketSpVtwoAcceptRequest() *CainiaoIotTicketSpVtwoAcceptAPIRequest{
    return &CainiaoIotTicketSpVtwoAcceptAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoIotTicketSpVtwoAcceptAPIRequest) GetApiMethodName() string {
    return "cainiao.iot.ticket.sp.vtwo.accept"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoIotTicketSpVtwoAcceptAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AcceptTicketTopRequest Setter
// 受理接口请求参数
func (r *CainiaoIotTicketSpVtwoAcceptAPIRequest) SetAcceptTicketTopRequest(_acceptTicketTopRequest *AcceptTicketV2TopRequest) error {
    r._acceptTicketTopRequest = _acceptTicketTopRequest
    r.Set("accept_ticket_top_request", _acceptTicketTopRequest)
    return nil
}

// AcceptTicketTopRequest Getter
func (r CainiaoIotTicketSpVtwoAcceptAPIRequest) GetAcceptTicketTopRequest() *AcceptTicketV2TopRequest {
    return r._acceptTicketTopRequest
}
