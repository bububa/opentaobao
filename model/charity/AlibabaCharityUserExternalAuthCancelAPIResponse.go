package charity

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCharityUserExternalAuthCancelAPIResponse 取消外部授权 API返回值
// alibaba.charity.user.external.auth.cancel
//
// 取消外部授权
type AlibabaCharityUserExternalAuthCancelAPIResponse struct {
	model.CommonResponse
	AlibabaCharityUserExternalAuthCancelAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCharityUserExternalAuthCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCharityUserExternalAuthCancelAPIResponseModel).Reset()
}

// AlibabaCharityUserExternalAuthCancelAPIResponseModel is 取消外部授权 成功返回结果
type AlibabaCharityUserExternalAuthCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_charity_user_external_auth_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *ThreehoursResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCharityUserExternalAuthCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCharityUserExternalAuthCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCharityUserExternalAuthCancelAPIResponse)
	},
}

// GetAlibabaCharityUserExternalAuthCancelAPIResponse 从 sync.Pool 获取 AlibabaCharityUserExternalAuthCancelAPIResponse
func GetAlibabaCharityUserExternalAuthCancelAPIResponse() *AlibabaCharityUserExternalAuthCancelAPIResponse {
	return poolAlibabaCharityUserExternalAuthCancelAPIResponse.Get().(*AlibabaCharityUserExternalAuthCancelAPIResponse)
}

// ReleaseAlibabaCharityUserExternalAuthCancelAPIResponse 将 AlibabaCharityUserExternalAuthCancelAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCharityUserExternalAuthCancelAPIResponse(v *AlibabaCharityUserExternalAuthCancelAPIResponse) {
	v.Reset()
	poolAlibabaCharityUserExternalAuthCancelAPIResponse.Put(v)
}
