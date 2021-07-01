package iotticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
IoT售后服务商确认接单 API返回值 
cainiao.iot.ticket.sp.vtwo.accept

IoT售后服务商确认接单
*/
type CainiaoIotTicketSpVtwoAcceptAPIResponse struct {
    model.CommonResponse
    CainiaoIotTicketSpVtwoAcceptResponse
}

// IoT售后服务商确认接单 成功返回结果
type CainiaoIotTicketSpVtwoAcceptResponse struct {
    XMLName xml.Name `xml:"cainiao_iot_ticket_sp_vtwo_accept_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *CainiaoIotTicketSpVtwoAcceptResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
