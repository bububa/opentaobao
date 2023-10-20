package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIResponse 是否支持云回看新SDK API返回值
// alibaba.ailabs.tmallgenie.sdk.device.issupportsdk
//
// 是否支持云回看新SDK
type AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIResponseModel).Reset()
}

// AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIResponseModel is 是否支持云回看新SDK 成功返回结果
type AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_sdk_device_issupportsdk_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAilabsTmallgenieSdkDeviceIssupportsdkResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIResponse)
	},
}

// GetAlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIResponse 从 sync.Pool 获取 AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIResponse
func GetAlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIResponse() *AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIResponse {
	return poolAlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIResponse.Get().(*AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIResponse)
}

// ReleaseAlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIResponse 将 AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIResponse(v *AlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIResponse) {
	v.Reset()
	poolAlibabaAilabsTmallgenieSdkDeviceIssupportsdkAPIResponse.Put(v)
}
