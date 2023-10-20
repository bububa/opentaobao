package maitix

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixEticketDistributionQueryAPIResponse 分销电子票查询接口 API返回值
// alibaba.damai.maitix.eticket.distribution.query
//
// 分销电子票查询接口
type AlibabaDamaiMaitixEticketDistributionQueryAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMaitixEticketDistributionQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixEticketDistributionQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMaitixEticketDistributionQueryAPIResponseModel).Reset()
}

// AlibabaDamaiMaitixEticketDistributionQueryAPIResponseModel is 分销电子票查询接口 成功返回结果
type AlibabaDamaiMaitixEticketDistributionQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_maitix_eticket_distribution_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *OpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixEticketDistributionQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMaitixEticketDistributionQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMaitixEticketDistributionQueryAPIResponse)
	},
}

// GetAlibabaDamaiMaitixEticketDistributionQueryAPIResponse 从 sync.Pool 获取 AlibabaDamaiMaitixEticketDistributionQueryAPIResponse
func GetAlibabaDamaiMaitixEticketDistributionQueryAPIResponse() *AlibabaDamaiMaitixEticketDistributionQueryAPIResponse {
	return poolAlibabaDamaiMaitixEticketDistributionQueryAPIResponse.Get().(*AlibabaDamaiMaitixEticketDistributionQueryAPIResponse)
}

// ReleaseAlibabaDamaiMaitixEticketDistributionQueryAPIResponse 将 AlibabaDamaiMaitixEticketDistributionQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMaitixEticketDistributionQueryAPIResponse(v *AlibabaDamaiMaitixEticketDistributionQueryAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMaitixEticketDistributionQueryAPIResponse.Put(v)
}
