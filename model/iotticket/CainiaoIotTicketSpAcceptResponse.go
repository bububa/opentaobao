package iotticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
IoT售后服务商确认接单 APIResponse
cainiao.iot.ticket.sp.accept

IoT售后服务商确认接单
*/
type CainiaoIotTicketSpAcceptAPIResponse struct {
    model.CommonResponse
    CainiaoIotTicketSpAcceptResponse
}

type CainiaoIotTicketSpAcceptResponse struct {
    XMLName xml.Name `xml:"cainiao_iot_ticket_sp_accept_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *CainiaoIotTicketSpAcceptResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
