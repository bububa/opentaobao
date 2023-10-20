package maitix

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixProjectDistributionQuerybypageAPIResponse 分销项目分页查询项目列表服务 API返回值
// alibaba.damai.maitix.project.distribution.querybypage
//
// 分销项目分页查询项目列表服务
type AlibabaDamaiMaitixProjectDistributionQuerybypageAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMaitixProjectDistributionQuerybypageAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixProjectDistributionQuerybypageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMaitixProjectDistributionQuerybypageAPIResponseModel).Reset()
}

// AlibabaDamaiMaitixProjectDistributionQuerybypageAPIResponseModel is 分销项目分页查询项目列表服务 成功返回结果
type AlibabaDamaiMaitixProjectDistributionQuerybypageAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_maitix_project_distribution_querybypage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *OpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixProjectDistributionQuerybypageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMaitixProjectDistributionQuerybypageAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMaitixProjectDistributionQuerybypageAPIResponse)
	},
}

// GetAlibabaDamaiMaitixProjectDistributionQuerybypageAPIResponse 从 sync.Pool 获取 AlibabaDamaiMaitixProjectDistributionQuerybypageAPIResponse
func GetAlibabaDamaiMaitixProjectDistributionQuerybypageAPIResponse() *AlibabaDamaiMaitixProjectDistributionQuerybypageAPIResponse {
	return poolAlibabaDamaiMaitixProjectDistributionQuerybypageAPIResponse.Get().(*AlibabaDamaiMaitixProjectDistributionQuerybypageAPIResponse)
}

// ReleaseAlibabaDamaiMaitixProjectDistributionQuerybypageAPIResponse 将 AlibabaDamaiMaitixProjectDistributionQuerybypageAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMaitixProjectDistributionQuerybypageAPIResponse(v *AlibabaDamaiMaitixProjectDistributionQuerybypageAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMaitixProjectDistributionQuerybypageAPIResponse.Put(v)
}
