package maitix

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIResponse 分销状态查询接口queryPerformStatusByPerformId API返回值
// alibaba.damai.maitix.opengateway.perform.status.query
//
// queryPerformStatusByPerformId
type AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIResponseModel).Reset()
}

// AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIResponseModel is 分销状态查询接口queryPerformStatusByPerformId 成功返回结果
type AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_maitix_opengateway_perform_status_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果
	Result *OpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIResponse)
	},
}

// GetAlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIResponse 从 sync.Pool 获取 AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIResponse
func GetAlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIResponse() *AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIResponse {
	return poolAlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIResponse.Get().(*AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIResponse)
}

// ReleaseAlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIResponse 将 AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIResponse(v *AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIResponse.Put(v)
}
