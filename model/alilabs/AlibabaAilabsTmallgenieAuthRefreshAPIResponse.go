package alilabs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieAuthRefreshAPIResponse 刷新token API返回值
// alibaba.ailabs.tmallgenie.auth.refresh
//
// 通过此接口刷新天猫精灵授权token
type AlibabaAilabsTmallgenieAuthRefreshAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsTmallgenieAuthRefreshAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsTmallgenieAuthRefreshAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsTmallgenieAuthRefreshAPIResponseModel).Reset()
}

// AlibabaAilabsTmallgenieAuthRefreshAPIResponseModel is 刷新token 成功返回结果
type AlibabaAilabsTmallgenieAuthRefreshAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_refresh_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// access_token
	AccessToken string `json:"access_token,omitempty" xml:"access_token,omitempty"`
	// refresh_token
	RefreshToken string `json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
	// access token过期时间，相对时间，单位为秒
	AccessExpiresIn int64 `json:"access_expires_in,omitempty" xml:"access_expires_in,omitempty"`
	// refresh token过期时间，相对时间，单位为秒
	RefreshExpiresIn int64 `json:"refresh_expires_in,omitempty" xml:"refresh_expires_in,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsTmallgenieAuthRefreshAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AccessToken = ""
	m.RefreshToken = ""
	m.AccessExpiresIn = 0
	m.RefreshExpiresIn = 0
}

var poolAlibabaAilabsTmallgenieAuthRefreshAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsTmallgenieAuthRefreshAPIResponse)
	},
}

// GetAlibabaAilabsTmallgenieAuthRefreshAPIResponse 从 sync.Pool 获取 AlibabaAilabsTmallgenieAuthRefreshAPIResponse
func GetAlibabaAilabsTmallgenieAuthRefreshAPIResponse() *AlibabaAilabsTmallgenieAuthRefreshAPIResponse {
	return poolAlibabaAilabsTmallgenieAuthRefreshAPIResponse.Get().(*AlibabaAilabsTmallgenieAuthRefreshAPIResponse)
}

// ReleaseAlibabaAilabsTmallgenieAuthRefreshAPIResponse 将 AlibabaAilabsTmallgenieAuthRefreshAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsTmallgenieAuthRefreshAPIResponse(v *AlibabaAilabsTmallgenieAuthRefreshAPIResponse) {
	v.Reset()
	poolAlibabaAilabsTmallgenieAuthRefreshAPIResponse.Put(v)
}
