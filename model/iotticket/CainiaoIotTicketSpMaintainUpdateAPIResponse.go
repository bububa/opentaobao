package iotticket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoIotTicketSpMaintainUpdateAPIResponse IoT售后服务商维修方案更新 API返回值
// cainiao.iot.ticket.sp.maintain.update
//
// IoT售后服务商维修方案更新
type CainiaoIotTicketSpMaintainUpdateAPIResponse struct {
	model.CommonResponse
	CainiaoIotTicketSpMaintainUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoIotTicketSpMaintainUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoIotTicketSpMaintainUpdateAPIResponseModel).Reset()
}

// CainiaoIotTicketSpMaintainUpdateAPIResponseModel is IoT售后服务商维修方案更新 成功返回结果
type CainiaoIotTicketSpMaintainUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_iot_ticket_sp_maintain_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *CainiaoIotTicketSpMaintainUpdateResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoIotTicketSpMaintainUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoIotTicketSpMaintainUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoIotTicketSpMaintainUpdateAPIResponse)
	},
}

// GetCainiaoIotTicketSpMaintainUpdateAPIResponse 从 sync.Pool 获取 CainiaoIotTicketSpMaintainUpdateAPIResponse
func GetCainiaoIotTicketSpMaintainUpdateAPIResponse() *CainiaoIotTicketSpMaintainUpdateAPIResponse {
	return poolCainiaoIotTicketSpMaintainUpdateAPIResponse.Get().(*CainiaoIotTicketSpMaintainUpdateAPIResponse)
}

// ReleaseCainiaoIotTicketSpMaintainUpdateAPIResponse 将 CainiaoIotTicketSpMaintainUpdateAPIResponse 保存到 sync.Pool
func ReleaseCainiaoIotTicketSpMaintainUpdateAPIResponse(v *CainiaoIotTicketSpMaintainUpdateAPIResponse) {
	v.Reset()
	poolCainiaoIotTicketSpMaintainUpdateAPIResponse.Put(v)
}
