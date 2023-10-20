package maitix

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixDistributionExchangepointQueryAPIResponse 分销查询取票点接口 API返回值
// alibaba.damai.maitix.distribution.exchangepoint.query
//
// 分销查询取票点接口
type AlibabaDamaiMaitixDistributionExchangepointQueryAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMaitixDistributionExchangepointQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixDistributionExchangepointQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMaitixDistributionExchangepointQueryAPIResponseModel).Reset()
}

// AlibabaDamaiMaitixDistributionExchangepointQueryAPIResponseModel is 分销查询取票点接口 成功返回结果
type AlibabaDamaiMaitixDistributionExchangepointQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_maitix_distribution_exchangepoint_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *OpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixDistributionExchangepointQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMaitixDistributionExchangepointQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMaitixDistributionExchangepointQueryAPIResponse)
	},
}

// GetAlibabaDamaiMaitixDistributionExchangepointQueryAPIResponse 从 sync.Pool 获取 AlibabaDamaiMaitixDistributionExchangepointQueryAPIResponse
func GetAlibabaDamaiMaitixDistributionExchangepointQueryAPIResponse() *AlibabaDamaiMaitixDistributionExchangepointQueryAPIResponse {
	return poolAlibabaDamaiMaitixDistributionExchangepointQueryAPIResponse.Get().(*AlibabaDamaiMaitixDistributionExchangepointQueryAPIResponse)
}

// ReleaseAlibabaDamaiMaitixDistributionExchangepointQueryAPIResponse 将 AlibabaDamaiMaitixDistributionExchangepointQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMaitixDistributionExchangepointQueryAPIResponse(v *AlibabaDamaiMaitixDistributionExchangepointQueryAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMaitixDistributionExchangepointQueryAPIResponse.Put(v)
}
