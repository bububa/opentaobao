package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectQueryAPIResponse 查询楼盘相关信息 API返回值
// alibaba.alihouse.newhome.project.query
//
// 根据outerid查询楼盘相关信息
type AlibabaAlihouseNewhomeProjectQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeProjectQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeProjectQueryAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeProjectQueryAPIResponseModel is 查询楼盘相关信息 成功返回结果
type AlibabaAlihouseNewhomeProjectQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeProjectQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeProjectQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectQueryAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeProjectQueryAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectQueryAPIResponse
func GetAlibabaAlihouseNewhomeProjectQueryAPIResponse() *AlibabaAlihouseNewhomeProjectQueryAPIResponse {
	return poolAlibabaAlihouseNewhomeProjectQueryAPIResponse.Get().(*AlibabaAlihouseNewhomeProjectQueryAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeProjectQueryAPIResponse 将 AlibabaAlihouseNewhomeProjectQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectQueryAPIResponse(v *AlibabaAlihouseNewhomeProjectQueryAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectQueryAPIResponse.Put(v)
}
