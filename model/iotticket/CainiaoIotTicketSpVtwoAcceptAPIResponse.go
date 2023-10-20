package iotticket

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoiotticketspvtwoacceptAPIResponse IoT售后服务商确认接单 API返回值
// cainiao.iot.ticket.sp.vtwo.accept
//
// IoT售后服务商确认接单
type CainiaoiotticketspvtwoacceptAPIResponse struct {
	model.CommonResponse
	CainiaoiotticketspvtwoacceptAPIResponseModel
}

// CainiaoiotticketspvtwoacceptAPIResponseModel is IoT售后服务商确认接单 成功返回结果
type CainiaoiotticketspvtwoacceptAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_iot_ticket_sp_vtwo_accept_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *CainiaoiotticketspvtwoacceptResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
