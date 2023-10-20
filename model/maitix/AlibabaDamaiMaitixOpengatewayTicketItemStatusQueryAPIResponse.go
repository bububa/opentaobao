package maitix

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIResponse 分销状态查询接口queryTicketItemStatusByTicketItemId API返回值
// alibaba.damai.maitix.opengateway.ticketItem.status.query
//
// queryTicketItemStatusByTicketItemId
type AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIResponseModel).Reset()
}

// AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIResponseModel is 分销状态查询接口queryTicketItemStatusByTicketItemId 成功返回结果
type AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_maitix_opengateway_ticketItem_status_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *OpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIResponse)
	},
}

// GetAlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIResponse 从 sync.Pool 获取 AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIResponse
func GetAlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIResponse() *AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIResponse {
	return poolAlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIResponse.Get().(*AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIResponse)
}

// ReleaseAlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIResponse 将 AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIResponse(v *AlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMaitixOpengatewayTicketItemStatusQueryAPIResponse.Put(v)
}
