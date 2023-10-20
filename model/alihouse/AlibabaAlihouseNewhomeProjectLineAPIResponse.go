package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectLineAPIResponse 楼盘上下架 API返回值
// alibaba.alihouse.newhome.project.line
//
// 上下架楼盘
type AlibabaAlihouseNewhomeProjectLineAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeProjectLineAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectLineAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeProjectLineAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeProjectLineAPIResponseModel is 楼盘上下架 成功返回结果
type AlibabaAlihouseNewhomeProjectLineAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_line_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeProjectLineResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectLineAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeProjectLineAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectLineAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeProjectLineAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectLineAPIResponse
func GetAlibabaAlihouseNewhomeProjectLineAPIResponse() *AlibabaAlihouseNewhomeProjectLineAPIResponse {
	return poolAlibabaAlihouseNewhomeProjectLineAPIResponse.Get().(*AlibabaAlihouseNewhomeProjectLineAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeProjectLineAPIResponse 将 AlibabaAlihouseNewhomeProjectLineAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectLineAPIResponse(v *AlibabaAlihouseNewhomeProjectLineAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectLineAPIResponse.Put(v)
}
