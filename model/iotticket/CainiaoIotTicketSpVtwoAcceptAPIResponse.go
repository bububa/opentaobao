package iotticket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoIotTicketSpVtwoAcceptAPIResponse IoT售后服务商确认接单 API返回值
// cainiao.iot.ticket.sp.vtwo.accept
//
// IoT售后服务商确认接单
type CainiaoIotTicketSpVtwoAcceptAPIResponse struct {
	model.CommonResponse
	CainiaoIotTicketSpVtwoAcceptAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoIotTicketSpVtwoAcceptAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoIotTicketSpVtwoAcceptAPIResponseModel).Reset()
}

// CainiaoIotTicketSpVtwoAcceptAPIResponseModel is IoT售后服务商确认接单 成功返回结果
type CainiaoIotTicketSpVtwoAcceptAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_iot_ticket_sp_vtwo_accept_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *CainiaoIotTicketSpVtwoAcceptResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoIotTicketSpVtwoAcceptAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoIotTicketSpVtwoAcceptAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoIotTicketSpVtwoAcceptAPIResponse)
	},
}

// GetCainiaoIotTicketSpVtwoAcceptAPIResponse 从 sync.Pool 获取 CainiaoIotTicketSpVtwoAcceptAPIResponse
func GetCainiaoIotTicketSpVtwoAcceptAPIResponse() *CainiaoIotTicketSpVtwoAcceptAPIResponse {
	return poolCainiaoIotTicketSpVtwoAcceptAPIResponse.Get().(*CainiaoIotTicketSpVtwoAcceptAPIResponse)
}

// ReleaseCainiaoIotTicketSpVtwoAcceptAPIResponse 将 CainiaoIotTicketSpVtwoAcceptAPIResponse 保存到 sync.Pool
func ReleaseCainiaoIotTicketSpVtwoAcceptAPIResponse(v *CainiaoIotTicketSpVtwoAcceptAPIResponse) {
	v.Reset()
	poolCainiaoIotTicketSpVtwoAcceptAPIResponse.Put(v)
}
