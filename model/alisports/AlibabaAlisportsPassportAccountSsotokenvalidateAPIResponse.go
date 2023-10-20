package alisports

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsPassportAccountSsotokenvalidateAPIResponse sso_token验证 API返回值
// alibaba.alisports.passport.account.ssotokenvalidate
//
// sso_token验证
type AlibabaAlisportsPassportAccountSsotokenvalidateAPIResponse struct {
	model.CommonResponse
	AlibabaAlisportsPassportAccountSsotokenvalidateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlisportsPassportAccountSsotokenvalidateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlisportsPassportAccountSsotokenvalidateAPIResponseModel).Reset()
}

// AlibabaAlisportsPassportAccountSsotokenvalidateAPIResponseModel is sso_token验证 成功返回结果
type AlibabaAlisportsPassportAccountSsotokenvalidateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alisports_passport_account_ssotokenvalidate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 提示信息
	AlispMsg string `json:"alisp_msg,omitempty" xml:"alisp_msg,omitempty"`
	// 阿里体育用户ID
	Aliuid string `json:"aliuid,omitempty" xml:"aliuid,omitempty"`
	// 第三方用户ID
	Appuid string `json:"appuid,omitempty" xml:"appuid,omitempty"`
	// 状态码 200表示操作成功
	AlispCode int64 `json:"alisp_code,omitempty" xml:"alisp_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlisportsPassportAccountSsotokenvalidateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AlispMsg = ""
	m.Aliuid = ""
	m.Appuid = ""
	m.AlispCode = 0
}

var poolAlibabaAlisportsPassportAccountSsotokenvalidateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlisportsPassportAccountSsotokenvalidateAPIResponse)
	},
}

// GetAlibabaAlisportsPassportAccountSsotokenvalidateAPIResponse 从 sync.Pool 获取 AlibabaAlisportsPassportAccountSsotokenvalidateAPIResponse
func GetAlibabaAlisportsPassportAccountSsotokenvalidateAPIResponse() *AlibabaAlisportsPassportAccountSsotokenvalidateAPIResponse {
	return poolAlibabaAlisportsPassportAccountSsotokenvalidateAPIResponse.Get().(*AlibabaAlisportsPassportAccountSsotokenvalidateAPIResponse)
}

// ReleaseAlibabaAlisportsPassportAccountSsotokenvalidateAPIResponse 将 AlibabaAlisportsPassportAccountSsotokenvalidateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlisportsPassportAccountSsotokenvalidateAPIResponse(v *AlibabaAlisportsPassportAccountSsotokenvalidateAPIResponse) {
	v.Reset()
	poolAlibabaAlisportsPassportAccountSsotokenvalidateAPIResponse.Put(v)
}
