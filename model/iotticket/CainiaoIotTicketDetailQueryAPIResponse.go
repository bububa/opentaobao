package iotticket

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoIotTicketDetailQueryAPIResponse IoT售后工单详情查询 API返回值
// cainiao.iot.ticket.detail.query
//
// Iot售后工单详情信息查询
type CainiaoIotTicketDetailQueryAPIResponse struct {
	model.CommonResponse
	CainiaoIotTicketDetailQueryAPIResponseModel
}

// CainiaoIotTicketDetailQueryAPIResponseModel is IoT售后工单详情查询 成功返回结果
type CainiaoIotTicketDetailQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_iot_ticket_detail_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *CainiaoIotTicketDetailQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
