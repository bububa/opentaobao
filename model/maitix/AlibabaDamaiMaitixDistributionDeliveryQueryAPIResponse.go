package maitix

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixDistributionDeliveryQueryAPIResponse 查询分销物流单 API返回值
// alibaba.damai.maitix.distribution.delivery.query
//
// 渠道查询物流订单
type AlibabaDamaiMaitixDistributionDeliveryQueryAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMaitixDistributionDeliveryQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixDistributionDeliveryQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMaitixDistributionDeliveryQueryAPIResponseModel).Reset()
}

// AlibabaDamaiMaitixDistributionDeliveryQueryAPIResponseModel is 查询分销物流单 成功返回结果
type AlibabaDamaiMaitixDistributionDeliveryQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_maitix_distribution_delivery_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *OpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixDistributionDeliveryQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMaitixDistributionDeliveryQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMaitixDistributionDeliveryQueryAPIResponse)
	},
}

// GetAlibabaDamaiMaitixDistributionDeliveryQueryAPIResponse 从 sync.Pool 获取 AlibabaDamaiMaitixDistributionDeliveryQueryAPIResponse
func GetAlibabaDamaiMaitixDistributionDeliveryQueryAPIResponse() *AlibabaDamaiMaitixDistributionDeliveryQueryAPIResponse {
	return poolAlibabaDamaiMaitixDistributionDeliveryQueryAPIResponse.Get().(*AlibabaDamaiMaitixDistributionDeliveryQueryAPIResponse)
}

// ReleaseAlibabaDamaiMaitixDistributionDeliveryQueryAPIResponse 将 AlibabaDamaiMaitixDistributionDeliveryQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMaitixDistributionDeliveryQueryAPIResponse(v *AlibabaDamaiMaitixDistributionDeliveryQueryAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMaitixDistributionDeliveryQueryAPIResponse.Put(v)
}
