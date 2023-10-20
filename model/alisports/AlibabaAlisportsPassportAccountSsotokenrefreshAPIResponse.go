package alisports

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsPassportAccountSsotokenrefreshAPIResponse sso_token刷新 API返回值
// alibaba.alisports.passport.account.ssotokenrefresh
//
// sso_token刷新
type AlibabaAlisportsPassportAccountSsotokenrefreshAPIResponse struct {
	model.CommonResponse
	AlibabaAlisportsPassportAccountSsotokenrefreshAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlisportsPassportAccountSsotokenrefreshAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlisportsPassportAccountSsotokenrefreshAPIResponseModel).Reset()
}

// AlibabaAlisportsPassportAccountSsotokenrefreshAPIResponseModel is sso_token刷新 成功返回结果
type AlibabaAlisportsPassportAccountSsotokenrefreshAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alisports_passport_account_ssotokenrefresh_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 状态码 200表示操作成功
	AlispCode string `json:"alisp_code,omitempty" xml:"alisp_code,omitempty"`
	// 提示信息
	AlispMsg string `json:"alisp_msg,omitempty" xml:"alisp_msg,omitempty"`
	// alisp_data
	AlispData *AlispData `json:"alisp_data,omitempty" xml:"alisp_data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlisportsPassportAccountSsotokenrefreshAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AlispCode = ""
	m.AlispMsg = ""
	m.AlispData = nil
}

var poolAlibabaAlisportsPassportAccountSsotokenrefreshAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlisportsPassportAccountSsotokenrefreshAPIResponse)
	},
}

// GetAlibabaAlisportsPassportAccountSsotokenrefreshAPIResponse 从 sync.Pool 获取 AlibabaAlisportsPassportAccountSsotokenrefreshAPIResponse
func GetAlibabaAlisportsPassportAccountSsotokenrefreshAPIResponse() *AlibabaAlisportsPassportAccountSsotokenrefreshAPIResponse {
	return poolAlibabaAlisportsPassportAccountSsotokenrefreshAPIResponse.Get().(*AlibabaAlisportsPassportAccountSsotokenrefreshAPIResponse)
}

// ReleaseAlibabaAlisportsPassportAccountSsotokenrefreshAPIResponse 将 AlibabaAlisportsPassportAccountSsotokenrefreshAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlisportsPassportAccountSsotokenrefreshAPIResponse(v *AlibabaAlisportsPassportAccountSsotokenrefreshAPIResponse) {
	v.Reset()
	poolAlibabaAlisportsPassportAccountSsotokenrefreshAPIResponse.Put(v)
}
