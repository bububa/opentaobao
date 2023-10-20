package charity

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIResponse 查询是否绑定3小时账号 API返回值
// alibaba.charity.charitytime.user.querythirduserhasauth
//
// 查询是否绑定3小时账号
type AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIResponse struct {
	model.CommonResponse
	AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIResponseModel).Reset()
}

// AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIResponseModel is 查询是否绑定3小时账号 成功返回结果
type AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_charity_charitytime_user_querythirduserhasauth_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *CsrResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCharityCharitytimeUserQuerythirduserhasauthAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIResponse)
	},
}

// GetAlibabaCharityCharitytimeUserQuerythirduserhasauthAPIResponse 从 sync.Pool 获取 AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIResponse
func GetAlibabaCharityCharitytimeUserQuerythirduserhasauthAPIResponse() *AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIResponse {
	return poolAlibabaCharityCharitytimeUserQuerythirduserhasauthAPIResponse.Get().(*AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIResponse)
}

// ReleaseAlibabaCharityCharitytimeUserQuerythirduserhasauthAPIResponse 将 AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCharityCharitytimeUserQuerythirduserhasauthAPIResponse(v *AlibabaCharityCharitytimeUserQuerythirduserhasauthAPIResponse) {
	v.Reset()
	poolAlibabaCharityCharitytimeUserQuerythirduserhasauthAPIResponse.Put(v)
}
