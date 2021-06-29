package iotticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商制定维修费方案 APIResponse
cainiao.iot.ticket.sp.maintain.vtwo.create

服务商制定维修费方案
*/
type CainiaoIotTicketSpMaintainVtwoCreateAPIResponse struct {
    model.CommonResponse
    CainiaoIotTicketSpMaintainVtwoCreateResponse
}

type CainiaoIotTicketSpMaintainVtwoCreateResponse struct {
    XMLName xml.Name `xml:"cainiao_iot_ticket_sp_maintain_vtwo_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *CainiaoIotTicketSpMaintainVtwoCreateResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
