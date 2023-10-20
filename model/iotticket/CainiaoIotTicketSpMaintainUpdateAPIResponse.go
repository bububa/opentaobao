package iotticket

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoiotticketspmaintainupdateAPIResponse IoT售后服务商维修方案更新 API返回值
// cainiao.iot.ticket.sp.maintain.update
//
// IoT售后服务商维修方案更新
type CainiaoiotticketspmaintainupdateAPIResponse struct {
	model.CommonResponse
	CainiaoiotticketspmaintainupdateAPIResponseModel
}

// CainiaoiotticketspmaintainupdateAPIResponseModel is IoT售后服务商维修方案更新 成功返回结果
type CainiaoiotticketspmaintainupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_iot_ticket_sp_maintain_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *CainiaoiotticketspmaintainupdateResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
