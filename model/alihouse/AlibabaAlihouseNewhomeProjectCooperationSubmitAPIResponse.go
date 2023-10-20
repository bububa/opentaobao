package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectCooperationSubmitAPIResponse 提交KA合作楼盘 API返回值
// alibaba.alihouse.newhome.project.cooperation.submit
//
// 提交KA合作楼盘
type AlibabaAlihouseNewhomeProjectCooperationSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeProjectCooperationSubmitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectCooperationSubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeProjectCooperationSubmitAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeProjectCooperationSubmitAPIResponseModel is 提交KA合作楼盘 成功返回结果
type AlibabaAlihouseNewhomeProjectCooperationSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_cooperation_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeProjectCooperationSubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectCooperationSubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeProjectCooperationSubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectCooperationSubmitAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeProjectCooperationSubmitAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectCooperationSubmitAPIResponse
func GetAlibabaAlihouseNewhomeProjectCooperationSubmitAPIResponse() *AlibabaAlihouseNewhomeProjectCooperationSubmitAPIResponse {
	return poolAlibabaAlihouseNewhomeProjectCooperationSubmitAPIResponse.Get().(*AlibabaAlihouseNewhomeProjectCooperationSubmitAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeProjectCooperationSubmitAPIResponse 将 AlibabaAlihouseNewhomeProjectCooperationSubmitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectCooperationSubmitAPIResponse(v *AlibabaAlihouseNewhomeProjectCooperationSubmitAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectCooperationSubmitAPIResponse.Put(v)
}
