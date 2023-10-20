package cainiaohandover

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoglobalsolutionserviceresourcequeryAPIResponse 查询解决方案服务资源列表 API返回值
// cainiao.global.solution.service.resource.query
//
// 返回直接解决方案的指定物流服务的可用资源列表
type CainiaoglobalsolutionserviceresourcequeryAPIResponse struct {
	model.CommonResponse
	CainiaoglobalsolutionserviceresourcequeryAPIResponseModel
}

// CainiaoglobalsolutionserviceresourcequeryAPIResponseModel is 查询解决方案服务资源列表 成功返回结果
type CainiaoglobalsolutionserviceresourcequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_solution_service_resource_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果
	Result *GlspResponse `json:"result,omitempty" xml:"result,omitempty"`
}
