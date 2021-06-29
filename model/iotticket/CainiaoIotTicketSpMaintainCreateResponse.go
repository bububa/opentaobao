package iotticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
IoT售后服务商制定维修方案 APIResponse
cainiao.iot.ticket.sp.maintain.create

IoT售后服务商制定维修方案
*/
type CainiaoIotTicketSpMaintainCreateAPIResponse struct {
    model.CommonResponse
    CainiaoIotTicketSpMaintainCreateResponse
}

type CainiaoIotTicketSpMaintainCreateResponse struct {
    XMLName xml.Name `xml:"cainiao_iot_ticket_sp_maintain_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *CainiaoIotTicketSpMaintainCreateResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
