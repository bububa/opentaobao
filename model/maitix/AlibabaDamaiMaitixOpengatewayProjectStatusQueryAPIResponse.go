package maitix

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIResponse 分销状态查询接口queryProjectStatusByProjectId API返回值
// alibaba.damai.maitix.opengateway.project.status.query
//
// queryProjectStatusByProjectId
type AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIResponseModel).Reset()
}

// AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIResponseModel is 分销状态查询接口queryProjectStatusByProjectId 成功返回结果
type AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_maitix_opengateway_project_status_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *OpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIResponse)
	},
}

// GetAlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIResponse 从 sync.Pool 获取 AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIResponse
func GetAlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIResponse() *AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIResponse {
	return poolAlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIResponse.Get().(*AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIResponse)
}

// ReleaseAlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIResponse 将 AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIResponse(v *AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIResponse.Put(v)
}
