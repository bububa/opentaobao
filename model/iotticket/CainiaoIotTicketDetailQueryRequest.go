package iotticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
IoT售后工单详情查询 API请求
cainiao.iot.ticket.detail.query

Iot售后工单详情信息查询
*/
type CainiaoIotTicketDetailQueryAPIRequest struct {
    model.Params
    // 服务商唯一编码
    _spCode   string
    // 工单Id
    _ticketId   int64
}

// 初始化CainiaoIotTicketDetailQueryAPIRequest对象
func NewCainiaoIotTicketDetailQueryRequest() *CainiaoIotTicketDetailQueryAPIRequest{
    return &CainiaoIotTicketDetailQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoIotTicketDetailQueryAPIRequest) GetApiMethodName() string {
    return "cainiao.iot.ticket.detail.query"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoIotTicketDetailQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SpCode Setter
// 服务商唯一编码
func (r *CainiaoIotTicketDetailQueryAPIRequest) SetSpCode(_spCode string) error {
    r._spCode = _spCode
    r.Set("sp_code", _spCode)
    return nil
}

// SpCode Getter
func (r CainiaoIotTicketDetailQueryAPIRequest) GetSpCode() string {
    return r._spCode
}
// TicketId Setter
// 工单Id
func (r *CainiaoIotTicketDetailQueryAPIRequest) SetTicketId(_ticketId int64) error {
    r._ticketId = _ticketId
    r.Set("ticket_id", _ticketId)
    return nil
}

// TicketId Getter
func (r CainiaoIotTicketDetailQueryAPIRequest) GetTicketId() int64 {
    return r._ticketId
}
