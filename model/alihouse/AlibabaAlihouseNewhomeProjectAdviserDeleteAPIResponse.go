package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectAdviserDeleteAPIResponse 删除楼盘顾问 API返回值
// alibaba.alihouse.newhome.project.adviser.delete
//
// 删除楼盘顾问
type AlibabaAlihouseNewhomeProjectAdviserDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeProjectAdviserDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectAdviserDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeProjectAdviserDeleteAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeProjectAdviserDeleteAPIResponseModel is 删除楼盘顾问 成功返回结果
type AlibabaAlihouseNewhomeProjectAdviserDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_adviser_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeProjectAdviserDeleteResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectAdviserDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeProjectAdviserDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectAdviserDeleteAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeProjectAdviserDeleteAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectAdviserDeleteAPIResponse
func GetAlibabaAlihouseNewhomeProjectAdviserDeleteAPIResponse() *AlibabaAlihouseNewhomeProjectAdviserDeleteAPIResponse {
	return poolAlibabaAlihouseNewhomeProjectAdviserDeleteAPIResponse.Get().(*AlibabaAlihouseNewhomeProjectAdviserDeleteAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeProjectAdviserDeleteAPIResponse 将 AlibabaAlihouseNewhomeProjectAdviserDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectAdviserDeleteAPIResponse(v *AlibabaAlihouseNewhomeProjectAdviserDeleteAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectAdviserDeleteAPIResponse.Put(v)
}
