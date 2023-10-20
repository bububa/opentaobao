package alilabs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieAuthGetcodeAPIResponse 获取token API返回值
// alibaba.ailabs.tmallgenie.auth.getcode
//
// 获取天猫精灵authCode
type AlibabaAilabsTmallgenieAuthGetcodeAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsTmallgenieAuthGetcodeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsTmallgenieAuthGetcodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsTmallgenieAuthGetcodeAPIResponseModel).Reset()
}

// AlibabaAilabsTmallgenieAuthGetcodeAPIResponseModel is 获取token 成功返回结果
type AlibabaAilabsTmallgenieAuthGetcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_getcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 授权码 json 串，包含授权码和二维码URL
	AuthCode string `json:"auth_code,omitempty" xml:"auth_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsTmallgenieAuthGetcodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AuthCode = ""
}

var poolAlibabaAilabsTmallgenieAuthGetcodeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsTmallgenieAuthGetcodeAPIResponse)
	},
}

// GetAlibabaAilabsTmallgenieAuthGetcodeAPIResponse 从 sync.Pool 获取 AlibabaAilabsTmallgenieAuthGetcodeAPIResponse
func GetAlibabaAilabsTmallgenieAuthGetcodeAPIResponse() *AlibabaAilabsTmallgenieAuthGetcodeAPIResponse {
	return poolAlibabaAilabsTmallgenieAuthGetcodeAPIResponse.Get().(*AlibabaAilabsTmallgenieAuthGetcodeAPIResponse)
}

// ReleaseAlibabaAilabsTmallgenieAuthGetcodeAPIResponse 将 AlibabaAilabsTmallgenieAuthGetcodeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsTmallgenieAuthGetcodeAPIResponse(v *AlibabaAilabsTmallgenieAuthGetcodeAPIResponse) {
	v.Reset()
	poolAlibabaAilabsTmallgenieAuthGetcodeAPIResponse.Put(v)
}
