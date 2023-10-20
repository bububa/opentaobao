package maitix

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixDistributionDeliveryCalculateAPIResponse 计算渠道用户下单快递费 API返回值
// alibaba.damai.maitix.distribution.delivery.calculate
//
// 计算渠道用户下单快递费
type AlibabaDamaiMaitixDistributionDeliveryCalculateAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMaitixDistributionDeliveryCalculateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixDistributionDeliveryCalculateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMaitixDistributionDeliveryCalculateAPIResponseModel).Reset()
}

// AlibabaDamaiMaitixDistributionDeliveryCalculateAPIResponseModel is 计算渠道用户下单快递费 成功返回结果
type AlibabaDamaiMaitixDistributionDeliveryCalculateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_maitix_distribution_delivery_calculate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaDamaiMaitixDistributionDeliveryCalculateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixDistributionDeliveryCalculateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMaitixDistributionDeliveryCalculateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMaitixDistributionDeliveryCalculateAPIResponse)
	},
}

// GetAlibabaDamaiMaitixDistributionDeliveryCalculateAPIResponse 从 sync.Pool 获取 AlibabaDamaiMaitixDistributionDeliveryCalculateAPIResponse
func GetAlibabaDamaiMaitixDistributionDeliveryCalculateAPIResponse() *AlibabaDamaiMaitixDistributionDeliveryCalculateAPIResponse {
	return poolAlibabaDamaiMaitixDistributionDeliveryCalculateAPIResponse.Get().(*AlibabaDamaiMaitixDistributionDeliveryCalculateAPIResponse)
}

// ReleaseAlibabaDamaiMaitixDistributionDeliveryCalculateAPIResponse 将 AlibabaDamaiMaitixDistributionDeliveryCalculateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMaitixDistributionDeliveryCalculateAPIResponse(v *AlibabaDamaiMaitixDistributionDeliveryCalculateAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMaitixDistributionDeliveryCalculateAPIResponse.Put(v)
}
