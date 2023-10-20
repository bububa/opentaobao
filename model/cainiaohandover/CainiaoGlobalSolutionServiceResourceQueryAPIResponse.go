package cainiaohandover

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalSolutionServiceResourceQueryAPIResponse 查询解决方案服务资源列表 API返回值
// cainiao.global.solution.service.resource.query
//
// 返回直接解决方案的指定物流服务的可用资源列表
type CainiaoGlobalSolutionServiceResourceQueryAPIResponse struct {
	model.CommonResponse
	CainiaoGlobalSolutionServiceResourceQueryAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoGlobalSolutionServiceResourceQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoGlobalSolutionServiceResourceQueryAPIResponseModel).Reset()
}

// CainiaoGlobalSolutionServiceResourceQueryAPIResponseModel is 查询解决方案服务资源列表 成功返回结果
type CainiaoGlobalSolutionServiceResourceQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_solution_service_resource_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果
	Result *GlspResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoGlobalSolutionServiceResourceQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoGlobalSolutionServiceResourceQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoGlobalSolutionServiceResourceQueryAPIResponse)
	},
}

// GetCainiaoGlobalSolutionServiceResourceQueryAPIResponse 从 sync.Pool 获取 CainiaoGlobalSolutionServiceResourceQueryAPIResponse
func GetCainiaoGlobalSolutionServiceResourceQueryAPIResponse() *CainiaoGlobalSolutionServiceResourceQueryAPIResponse {
	return poolCainiaoGlobalSolutionServiceResourceQueryAPIResponse.Get().(*CainiaoGlobalSolutionServiceResourceQueryAPIResponse)
}

// ReleaseCainiaoGlobalSolutionServiceResourceQueryAPIResponse 将 CainiaoGlobalSolutionServiceResourceQueryAPIResponse 保存到 sync.Pool
func ReleaseCainiaoGlobalSolutionServiceResourceQueryAPIResponse(v *CainiaoGlobalSolutionServiceResourceQueryAPIResponse) {
	v.Reset()
	poolCainiaoGlobalSolutionServiceResourceQueryAPIResponse.Put(v)
}
