package iotticket

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoIotTicketSpAcceptAPIResponse
IoT售后服务商确认接单 API返回值
cainiao.iot.ticket.sp.accept

IoT售后服务商确认接单 */
type CainiaoIotTicketSpAcceptAPIResponse struct {
	model.CommonResponse
	CainiaoIotTicketSpAcceptAPIResponseModel
}

// CainiaoIotTicketSpAcceptAPIResponseModel is IoT售后服务商确认接单 成功返回结果
type CainiaoIotTicketSpAcceptAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_iot_ticket_sp_accept_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *CainiaoIotTicketSpAcceptResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
