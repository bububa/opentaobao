package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieAuthDeviceGetcodeAPIResponse 获取authcode API返回值
// alibaba.ailabs.tmallgenie.auth.device.getcode
//
// 获取绑定的authcode
type AlibabaAilabsTmallgenieAuthDeviceGetcodeAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsTmallgenieAuthDeviceGetcodeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsTmallgenieAuthDeviceGetcodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsTmallgenieAuthDeviceGetcodeAPIResponseModel).Reset()
}

// AlibabaAilabsTmallgenieAuthDeviceGetcodeAPIResponseModel is 获取authcode 成功返回结果
type AlibabaAilabsTmallgenieAuthDeviceGetcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_device_getcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// authcode
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsTmallgenieAuthDeviceGetcodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
	m.Message = ""
	m.ResultCode = 0
}

var poolAlibabaAilabsTmallgenieAuthDeviceGetcodeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsTmallgenieAuthDeviceGetcodeAPIResponse)
	},
}

// GetAlibabaAilabsTmallgenieAuthDeviceGetcodeAPIResponse 从 sync.Pool 获取 AlibabaAilabsTmallgenieAuthDeviceGetcodeAPIResponse
func GetAlibabaAilabsTmallgenieAuthDeviceGetcodeAPIResponse() *AlibabaAilabsTmallgenieAuthDeviceGetcodeAPIResponse {
	return poolAlibabaAilabsTmallgenieAuthDeviceGetcodeAPIResponse.Get().(*AlibabaAilabsTmallgenieAuthDeviceGetcodeAPIResponse)
}

// ReleaseAlibabaAilabsTmallgenieAuthDeviceGetcodeAPIResponse 将 AlibabaAilabsTmallgenieAuthDeviceGetcodeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsTmallgenieAuthDeviceGetcodeAPIResponse(v *AlibabaAilabsTmallgenieAuthDeviceGetcodeAPIResponse) {
	v.Reset()
	poolAlibabaAilabsTmallgenieAuthDeviceGetcodeAPIResponse.Put(v)
}
