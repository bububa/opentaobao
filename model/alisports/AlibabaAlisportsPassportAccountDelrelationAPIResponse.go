package alisports

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsPassportAccountDelrelationAPIResponse 阿里体育会员系统--取消三方关联接口 API返回值
// alibaba.alisports.passport.account.delrelation
//
// 阿里体育会员系统--取消三方关联接口
type AlibabaAlisportsPassportAccountDelrelationAPIResponse struct {
	model.CommonResponse
	AlibabaAlisportsPassportAccountDelrelationAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlisportsPassportAccountDelrelationAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlisportsPassportAccountDelrelationAPIResponseModel).Reset()
}

// AlibabaAlisportsPassportAccountDelrelationAPIResponseModel is 阿里体育会员系统--取消三方关联接口 成功返回结果
type AlibabaAlisportsPassportAccountDelrelationAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alisports_passport_account_delrelation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alisp_msg
	AlispMsg string `json:"alisp_msg,omitempty" xml:"alisp_msg,omitempty"`
	// alisp_data
	AlispData string `json:"alisp_data,omitempty" xml:"alisp_data,omitempty"`
	// alisp_code
	AlispCode int64 `json:"alisp_code,omitempty" xml:"alisp_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlisportsPassportAccountDelrelationAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AlispMsg = ""
	m.AlispData = ""
	m.AlispCode = 0
}

var poolAlibabaAlisportsPassportAccountDelrelationAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlisportsPassportAccountDelrelationAPIResponse)
	},
}

// GetAlibabaAlisportsPassportAccountDelrelationAPIResponse 从 sync.Pool 获取 AlibabaAlisportsPassportAccountDelrelationAPIResponse
func GetAlibabaAlisportsPassportAccountDelrelationAPIResponse() *AlibabaAlisportsPassportAccountDelrelationAPIResponse {
	return poolAlibabaAlisportsPassportAccountDelrelationAPIResponse.Get().(*AlibabaAlisportsPassportAccountDelrelationAPIResponse)
}

// ReleaseAlibabaAlisportsPassportAccountDelrelationAPIResponse 将 AlibabaAlisportsPassportAccountDelrelationAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlisportsPassportAccountDelrelationAPIResponse(v *AlibabaAlisportsPassportAccountDelrelationAPIResponse) {
	v.Reset()
	poolAlibabaAlisportsPassportAccountDelrelationAPIResponse.Put(v)
}
