package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectDynamicSubmitAPIResponse 提交楼盘快讯 API返回值
// alibaba.alihouse.newhome.project.dynamic.submit
//
// 提交楼盘快讯
type AlibabaAlihouseNewhomeProjectDynamicSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeProjectDynamicSubmitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectDynamicSubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeProjectDynamicSubmitAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeProjectDynamicSubmitAPIResponseModel is 提交楼盘快讯 成功返回结果
type AlibabaAlihouseNewhomeProjectDynamicSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_dynamic_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeProjectDynamicSubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectDynamicSubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeProjectDynamicSubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectDynamicSubmitAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeProjectDynamicSubmitAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectDynamicSubmitAPIResponse
func GetAlibabaAlihouseNewhomeProjectDynamicSubmitAPIResponse() *AlibabaAlihouseNewhomeProjectDynamicSubmitAPIResponse {
	return poolAlibabaAlihouseNewhomeProjectDynamicSubmitAPIResponse.Get().(*AlibabaAlihouseNewhomeProjectDynamicSubmitAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeProjectDynamicSubmitAPIResponse 将 AlibabaAlihouseNewhomeProjectDynamicSubmitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectDynamicSubmitAPIResponse(v *AlibabaAlihouseNewhomeProjectDynamicSubmitAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectDynamicSubmitAPIResponse.Put(v)
}
