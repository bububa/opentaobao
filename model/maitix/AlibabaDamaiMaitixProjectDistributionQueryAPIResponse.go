package maitix

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixProjectDistributionQueryAPIResponse 分销单个项目信息查询 API返回值
// alibaba.damai.maitix.project.distribution.query
//
// 发布分销项目查询单个项目信息接口
type AlibabaDamaiMaitixProjectDistributionQueryAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMaitixProjectDistributionQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixProjectDistributionQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMaitixProjectDistributionQueryAPIResponseModel).Reset()
}

// AlibabaDamaiMaitixProjectDistributionQueryAPIResponseModel is 分销单个项目信息查询 成功返回结果
type AlibabaDamaiMaitixProjectDistributionQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_maitix_project_distribution_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *OpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixProjectDistributionQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMaitixProjectDistributionQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMaitixProjectDistributionQueryAPIResponse)
	},
}

// GetAlibabaDamaiMaitixProjectDistributionQueryAPIResponse 从 sync.Pool 获取 AlibabaDamaiMaitixProjectDistributionQueryAPIResponse
func GetAlibabaDamaiMaitixProjectDistributionQueryAPIResponse() *AlibabaDamaiMaitixProjectDistributionQueryAPIResponse {
	return poolAlibabaDamaiMaitixProjectDistributionQueryAPIResponse.Get().(*AlibabaDamaiMaitixProjectDistributionQueryAPIResponse)
}

// ReleaseAlibabaDamaiMaitixProjectDistributionQueryAPIResponse 将 AlibabaDamaiMaitixProjectDistributionQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMaitixProjectDistributionQueryAPIResponse(v *AlibabaDamaiMaitixProjectDistributionQueryAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMaitixProjectDistributionQueryAPIResponse.Put(v)
}
