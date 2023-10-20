package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectPhoneSubmitAPIResponse 提交楼盘电话 API返回值
// alibaba.alihouse.newhome.project.phone.submit
//
// 提交楼盘电话
type AlibabaAlihouseNewhomeProjectPhoneSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeProjectPhoneSubmitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectPhoneSubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeProjectPhoneSubmitAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeProjectPhoneSubmitAPIResponseModel is 提交楼盘电话 成功返回结果
type AlibabaAlihouseNewhomeProjectPhoneSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_phone_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeProjectPhoneSubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectPhoneSubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeProjectPhoneSubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectPhoneSubmitAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeProjectPhoneSubmitAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectPhoneSubmitAPIResponse
func GetAlibabaAlihouseNewhomeProjectPhoneSubmitAPIResponse() *AlibabaAlihouseNewhomeProjectPhoneSubmitAPIResponse {
	return poolAlibabaAlihouseNewhomeProjectPhoneSubmitAPIResponse.Get().(*AlibabaAlihouseNewhomeProjectPhoneSubmitAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeProjectPhoneSubmitAPIResponse 将 AlibabaAlihouseNewhomeProjectPhoneSubmitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectPhoneSubmitAPIResponse(v *AlibabaAlihouseNewhomeProjectPhoneSubmitAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectPhoneSubmitAPIResponse.Put(v)
}
