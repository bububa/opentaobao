package maitix

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIResponse 查询招行支付状态api API返回值
// alibaba.damai.maitix.distribution.cmb.querypayresult
//
// queryPayResult
type AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIResponseModel).Reset()
}

// AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIResponseModel is 查询招行支付状态api 成功返回结果
type AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_maitix_distribution_cmb_querypayresult_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *OpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMaitixDistributionCmbQuerypayresultAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIResponse)
	},
}

// GetAlibabaDamaiMaitixDistributionCmbQuerypayresultAPIResponse 从 sync.Pool 获取 AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIResponse
func GetAlibabaDamaiMaitixDistributionCmbQuerypayresultAPIResponse() *AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIResponse {
	return poolAlibabaDamaiMaitixDistributionCmbQuerypayresultAPIResponse.Get().(*AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIResponse)
}

// ReleaseAlibabaDamaiMaitixDistributionCmbQuerypayresultAPIResponse 将 AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMaitixDistributionCmbQuerypayresultAPIResponse(v *AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMaitixDistributionCmbQuerypayresultAPIResponse.Put(v)
}
