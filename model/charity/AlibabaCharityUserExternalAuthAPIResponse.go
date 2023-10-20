package charity

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCharityUserExternalAuthAPIResponse 外部用户授权 API返回值
// alibaba.charity.user.external.auth
//
// 外部用户授权
type AlibabaCharityUserExternalAuthAPIResponse struct {
	model.CommonResponse
	AlibabaCharityUserExternalAuthAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCharityUserExternalAuthAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCharityUserExternalAuthAPIResponseModel).Reset()
}

// AlibabaCharityUserExternalAuthAPIResponseModel is 外部用户授权 成功返回结果
type AlibabaCharityUserExternalAuthAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_charity_user_external_auth_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *ThreehoursResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCharityUserExternalAuthAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCharityUserExternalAuthAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCharityUserExternalAuthAPIResponse)
	},
}

// GetAlibabaCharityUserExternalAuthAPIResponse 从 sync.Pool 获取 AlibabaCharityUserExternalAuthAPIResponse
func GetAlibabaCharityUserExternalAuthAPIResponse() *AlibabaCharityUserExternalAuthAPIResponse {
	return poolAlibabaCharityUserExternalAuthAPIResponse.Get().(*AlibabaCharityUserExternalAuthAPIResponse)
}

// ReleaseAlibabaCharityUserExternalAuthAPIResponse 将 AlibabaCharityUserExternalAuthAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCharityUserExternalAuthAPIResponse(v *AlibabaCharityUserExternalAuthAPIResponse) {
	v.Reset()
	poolAlibabaCharityUserExternalAuthAPIResponse.Put(v)
}
