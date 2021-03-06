package iotticket

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoIotTicketSpMaintainCreateAPIResponse IoT售后服务商制定维修方案 API返回值
// cainiao.iot.ticket.sp.maintain.create
//
// IoT售后服务商制定维修方案
type CainiaoIotTicketSpMaintainCreateAPIResponse struct {
	model.CommonResponse
	CainiaoIotTicketSpMaintainCreateAPIResponseModel
}

// CainiaoIotTicketSpMaintainCreateAPIResponseModel is IoT售后服务商制定维修方案 成功返回结果
type CainiaoIotTicketSpMaintainCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_iot_ticket_sp_maintain_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *CainiaoIotTicketSpMaintainCreateResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
