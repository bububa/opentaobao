package iotticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
IoT售后工单详情查询 APIRequest
cainiao.iot.ticket.detail.query

Iot售后工单详情信息查询
*/
type CainiaoIotTicketDetailQueryRequest struct {
    model.Params

    // 服务商唯一编码
    spCode   string 

    // 工单Id
    ticketId   int64 

}

func NewCainiaoIotTicketDetailQueryRequest() *CainiaoIotTicketDetailQueryRequest{
    return &CainiaoIotTicketDetailQueryRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoIotTicketDetailQueryRequest) GetApiMethodName() string {
    return "cainiao.iot.ticket.detail.query"
}

func (r CainiaoIotTicketDetailQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoIotTicketDetailQueryRequest) SetSpCode(spCode string) error {
    r.spCode = spCode
    r.Set("sp_code", spCode)
    return nil
}

func (r CainiaoIotTicketDetailQueryRequest) GetSpCode() string {
    return r.spCode
}

func (r *CainiaoIotTicketDetailQueryRequest) SetTicketId(ticketId int64) error {
    r.ticketId = ticketId
    r.Set("ticket_id", ticketId)
    return nil
}

func (r CainiaoIotTicketDetailQueryRequest) GetTicketId() int64 {
    return r.ticketId
}

