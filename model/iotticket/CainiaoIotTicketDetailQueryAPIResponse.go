package iotticket

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *CainiaoIotTicketDetailQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoIotTicketDetailQueryAPIResponseModel).Reset()
}

// CainiaoIotTicketDetailQueryAPIResponseModel is IoT售后工单详情查询 成功返回结果
type CainiaoIotTicketDetailQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_iot_ticket_detail_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *CainiaoIotTicketDetailQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoIotTicketDetailQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoIotTicketDetailQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoIotTicketDetailQueryAPIResponse)
	},
}

// GetCainiaoIotTicketDetailQueryAPIResponse 从 sync.Pool 获取 CainiaoIotTicketDetailQueryAPIResponse
func GetCainiaoIotTicketDetailQueryAPIResponse() *CainiaoIotTicketDetailQueryAPIResponse {
	return poolCainiaoIotTicketDetailQueryAPIResponse.Get().(*CainiaoIotTicketDetailQueryAPIResponse)
}

// ReleaseCainiaoIotTicketDetailQueryAPIResponse 将 CainiaoIotTicketDetailQueryAPIResponse 保存到 sync.Pool
func ReleaseCainiaoIotTicketDetailQueryAPIResponse(v *CainiaoIotTicketDetailQueryAPIResponse) {
	v.Reset()
	poolCainiaoIotTicketDetailQueryAPIResponse.Put(v)
}
