package maitix

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixProjectDistributionQuerylistAPIResponse 分销项目列表查询（已过时，不推荐使用） API返回值
// alibaba.damai.maitix.project.distribution.querylist
//
// 分销项目列表查询接口（已过时，不推荐使用）
type AlibabaDamaiMaitixProjectDistributionQuerylistAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMaitixProjectDistributionQuerylistAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixProjectDistributionQuerylistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMaitixProjectDistributionQuerylistAPIResponseModel).Reset()
}

// AlibabaDamaiMaitixProjectDistributionQuerylistAPIResponseModel is 分销项目列表查询（已过时，不推荐使用） 成功返回结果
type AlibabaDamaiMaitixProjectDistributionQuerylistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_maitix_project_distribution_querylist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *OpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixProjectDistributionQuerylistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMaitixProjectDistributionQuerylistAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMaitixProjectDistributionQuerylistAPIResponse)
	},
}

// GetAlibabaDamaiMaitixProjectDistributionQuerylistAPIResponse 从 sync.Pool 获取 AlibabaDamaiMaitixProjectDistributionQuerylistAPIResponse
func GetAlibabaDamaiMaitixProjectDistributionQuerylistAPIResponse() *AlibabaDamaiMaitixProjectDistributionQuerylistAPIResponse {
	return poolAlibabaDamaiMaitixProjectDistributionQuerylistAPIResponse.Get().(*AlibabaDamaiMaitixProjectDistributionQuerylistAPIResponse)
}

// ReleaseAlibabaDamaiMaitixProjectDistributionQuerylistAPIResponse 将 AlibabaDamaiMaitixProjectDistributionQuerylistAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMaitixProjectDistributionQuerylistAPIResponse(v *AlibabaDamaiMaitixProjectDistributionQuerylistAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMaitixProjectDistributionQuerylistAPIResponse.Put(v)
}
