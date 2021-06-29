package iotticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
IoT售后服务商维修方案更新 APIResponse
cainiao.iot.ticket.sp.maintain.update

IoT售后服务商维修方案更新
*/
type CainiaoIotTicketSpMaintainUpdateAPIResponse struct {
    model.CommonResponse
    CainiaoIotTicketSpMaintainUpdateResponse
}

type CainiaoIotTicketSpMaintainUpdateResponse struct {
    XMLName xml.Name `xml:"cainiao_iot_ticket_sp_maintain_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *CainiaoIotTicketSpMaintainUpdateResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
