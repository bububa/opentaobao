package iotticket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoIotTicketSpCommentAPIResponse IoT售后服务商工单备注 API返回值
// cainiao.iot.ticket.sp.comment
//
// IoT售后服务商工单备注
type CainiaoIotTicketSpCommentAPIResponse struct {
	model.CommonResponse
	CainiaoIotTicketSpCommentAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoIotTicketSpCommentAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoIotTicketSpCommentAPIResponseModel).Reset()
}

// CainiaoIotTicketSpCommentAPIResponseModel is IoT售后服务商工单备注 成功返回结果
type CainiaoIotTicketSpCommentAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_iot_ticket_sp_comment_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *CainiaoIotTicketSpCommentResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoIotTicketSpCommentAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoIotTicketSpCommentAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoIotTicketSpCommentAPIResponse)
	},
}

// GetCainiaoIotTicketSpCommentAPIResponse 从 sync.Pool 获取 CainiaoIotTicketSpCommentAPIResponse
func GetCainiaoIotTicketSpCommentAPIResponse() *CainiaoIotTicketSpCommentAPIResponse {
	return poolCainiaoIotTicketSpCommentAPIResponse.Get().(*CainiaoIotTicketSpCommentAPIResponse)
}

// ReleaseCainiaoIotTicketSpCommentAPIResponse 将 CainiaoIotTicketSpCommentAPIResponse 保存到 sync.Pool
func ReleaseCainiaoIotTicketSpCommentAPIResponse(v *CainiaoIotTicketSpCommentAPIResponse) {
	v.Reset()
	poolCainiaoIotTicketSpCommentAPIResponse.Put(v)
}
