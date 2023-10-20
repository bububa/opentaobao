package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangDeliveryDecisionQueryAPIResponse 查询黑白名单快递 API返回值
// alibaba.dchain.aoxiang.delivery.decision.query
//
// 查询黑白名单快递
type AlibabaDchainAoxiangDeliveryDecisionQueryAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangDeliveryDecisionQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangDeliveryDecisionQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangDeliveryDecisionQueryAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangDeliveryDecisionQueryAPIResponseModel is 查询黑白名单快递 成功返回结果
type AlibabaDchainAoxiangDeliveryDecisionQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_delivery_decision_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果
	DeliveryDecisionQueryResponse *DeliveryDecisionQueryResponse `json:"delivery_decision_query_response,omitempty" xml:"delivery_decision_query_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangDeliveryDecisionQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DeliveryDecisionQueryResponse = nil
}

var poolAlibabaDchainAoxiangDeliveryDecisionQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangDeliveryDecisionQueryAPIResponse)
	},
}

// GetAlibabaDchainAoxiangDeliveryDecisionQueryAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangDeliveryDecisionQueryAPIResponse
func GetAlibabaDchainAoxiangDeliveryDecisionQueryAPIResponse() *AlibabaDchainAoxiangDeliveryDecisionQueryAPIResponse {
	return poolAlibabaDchainAoxiangDeliveryDecisionQueryAPIResponse.Get().(*AlibabaDchainAoxiangDeliveryDecisionQueryAPIResponse)
}

// ReleaseAlibabaDchainAoxiangDeliveryDecisionQueryAPIResponse 将 AlibabaDchainAoxiangDeliveryDecisionQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangDeliveryDecisionQueryAPIResponse(v *AlibabaDchainAoxiangDeliveryDecisionQueryAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangDeliveryDecisionQueryAPIResponse.Put(v)
}
