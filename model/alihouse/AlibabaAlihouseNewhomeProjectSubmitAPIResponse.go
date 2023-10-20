package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectSubmitAPIResponse 提交楼盘信息 API返回值
// alibaba.alihouse.newhome.project.submit
//
// 提交楼盘信息
type AlibabaAlihouseNewhomeProjectSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeProjectSubmitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectSubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeProjectSubmitAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeProjectSubmitAPIResponseModel is 提交楼盘信息 成功返回结果
type AlibabaAlihouseNewhomeProjectSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回结果体
	Result *AlibabaAlihouseNewhomeProjectSubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectSubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeProjectSubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectSubmitAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeProjectSubmitAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectSubmitAPIResponse
func GetAlibabaAlihouseNewhomeProjectSubmitAPIResponse() *AlibabaAlihouseNewhomeProjectSubmitAPIResponse {
	return poolAlibabaAlihouseNewhomeProjectSubmitAPIResponse.Get().(*AlibabaAlihouseNewhomeProjectSubmitAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeProjectSubmitAPIResponse 将 AlibabaAlihouseNewhomeProjectSubmitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectSubmitAPIResponse(v *AlibabaAlihouseNewhomeProjectSubmitAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectSubmitAPIResponse.Put(v)
}
