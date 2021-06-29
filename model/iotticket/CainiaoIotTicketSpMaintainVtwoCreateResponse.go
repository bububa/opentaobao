package iotticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商制定维修费方案 API返回值 
cainiao.iot.ticket.sp.maintain.vtwo.create

服务商制定维修费方案
*/
type CainiaoIotTicketSpMaintainVtwoCreateAPIResponse struct {
    model.CommonResponse
    CainiaoIotTicketSpMaintainVtwoCreateResponse
}

// 服务商制定维修费方案 成功返回结果
type CainiaoIotTicketSpMaintainVtwoCreateResponse struct {
    XMLName xml.Name `xml:"cainiao_iot_ticket_sp_maintain_vtwo_create_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *CainiaoIotTicketSpMaintainVtwoCreateResultDTO `json:"result,omitempty" xml:"result,omitempty"`
}
