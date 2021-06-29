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
type CainiaoIotTicketDetailQueryRequest struct {
    model.Params
    // 服务商唯一编码
    _spCode   string
    // 工单Id
    _ticketId   int64
}

// 初始化CainiaoIotTicketDetailQueryRequest对象
func NewCainiaoIotTicketDetailQueryRequest() *CainiaoIotTicketDetailQueryRequest{
    return &CainiaoIotTicketDetailQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoIotTicketDetailQueryRequest) GetApiMethodName() string {
    return "cainiao.iot.ticket.detail.query"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoIotTicketDetailQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SpCode Setter
// 服务商唯一编码
func (r *CainiaoIotTicketDetailQueryRequest) SetSpCode(_spCode string) error {
    r._spCode = _spCode
    r.Set("sp_code", _spCode)
    return nil
}

// SpCode Getter
func (r CainiaoIotTicketDetailQueryRequest) GetSpCode() string {
    return r._spCode
}
// TicketId Setter
// 工单Id
func (r *CainiaoIotTicketDetailQueryRequest) SetTicketId(_ticketId int64) error {
    r._ticketId = _ticketId
    r.Set("ticket_id", _ticketId)
    return nil
}

// TicketId Getter
func (r CainiaoIotTicketDetailQueryRequest) GetTicketId() int64 {
    return r._ticketId
}
