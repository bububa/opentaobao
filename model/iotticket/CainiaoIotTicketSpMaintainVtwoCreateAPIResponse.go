package iotticket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoIotTicketSpMaintainVtwoCreateAPIResponse 服务商制定维修费方案 API返回值
// cainiao.iot.ticket.sp.maintain.vtwo.create
//
// 服务商制定维修费方案
type CainiaoIotTicketSpMaintainVtwoCreateAPIResponse struct {
	model.CommonResponse
	CainiaoIotTicketSpMaintainVtwoCreateAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoIotTicketSpMaintainVtwoCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoIotTicketSpMaintainVtwoCreateAPIResponseModel).Reset()
}

// CainiaoIotTicketSpMaintainVtwoCreateAPIResponseModel is 服务商制定维修费方案 成功返回结果
type CainiaoIotTicketSpMaintainVtwoCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_iot_ticket_sp_maintain_vtwo_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *CainiaoIotTicketSpMaintainVtwoCreateResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoIotTicketSpMaintainVtwoCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoIotTicketSpMaintainVtwoCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoIotTicketSpMaintainVtwoCreateAPIResponse)
	},
}

// GetCainiaoIotTicketSpMaintainVtwoCreateAPIResponse 从 sync.Pool 获取 CainiaoIotTicketSpMaintainVtwoCreateAPIResponse
func GetCainiaoIotTicketSpMaintainVtwoCreateAPIResponse() *CainiaoIotTicketSpMaintainVtwoCreateAPIResponse {
	return poolCainiaoIotTicketSpMaintainVtwoCreateAPIResponse.Get().(*CainiaoIotTicketSpMaintainVtwoCreateAPIResponse)
}

// ReleaseCainiaoIotTicketSpMaintainVtwoCreateAPIResponse 将 CainiaoIotTicketSpMaintainVtwoCreateAPIResponse 保存到 sync.Pool
func ReleaseCainiaoIotTicketSpMaintainVtwoCreateAPIResponse(v *CainiaoIotTicketSpMaintainVtwoCreateAPIResponse) {
	v.Reset()
	poolCainiaoIotTicketSpMaintainVtwoCreateAPIResponse.Put(v)
}
