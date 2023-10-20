package charity

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCharityCharitytimeUserCancelauthAPIResponse 取消用户授权 API返回值
// alibaba.charity.charitytime.user.cancelauth
//
// 取消用户授权
type AlibabaCharityCharitytimeUserCancelauthAPIResponse struct {
	model.CommonResponse
	AlibabaCharityCharitytimeUserCancelauthAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCharityCharitytimeUserCancelauthAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCharityCharitytimeUserCancelauthAPIResponseModel).Reset()
}

// AlibabaCharityCharitytimeUserCancelauthAPIResponseModel is 取消用户授权 成功返回结果
type AlibabaCharityCharitytimeUserCancelauthAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_charity_charitytime_user_cancelauth_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *CsrResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCharityCharitytimeUserCancelauthAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCharityCharitytimeUserCancelauthAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCharityCharitytimeUserCancelauthAPIResponse)
	},
}

// GetAlibabaCharityCharitytimeUserCancelauthAPIResponse 从 sync.Pool 获取 AlibabaCharityCharitytimeUserCancelauthAPIResponse
func GetAlibabaCharityCharitytimeUserCancelauthAPIResponse() *AlibabaCharityCharitytimeUserCancelauthAPIResponse {
	return poolAlibabaCharityCharitytimeUserCancelauthAPIResponse.Get().(*AlibabaCharityCharitytimeUserCancelauthAPIResponse)
}

// ReleaseAlibabaCharityCharitytimeUserCancelauthAPIResponse 将 AlibabaCharityCharitytimeUserCancelauthAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCharityCharitytimeUserCancelauthAPIResponse(v *AlibabaCharityCharitytimeUserCancelauthAPIResponse) {
	v.Reset()
	poolAlibabaCharityCharitytimeUserCancelauthAPIResponse.Put(v)
}
