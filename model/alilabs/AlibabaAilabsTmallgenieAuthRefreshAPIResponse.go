package alilabs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsTmallgenieAuthRefreshAPIResponse
刷新token API返回值
alibaba.ailabs.tmallgenie.auth.refresh

通过此接口刷新天猫精灵授权token */
type AlibabaAilabsTmallgenieAuthRefreshAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsTmallgenieAuthRefreshAPIResponseModel
}

// AlibabaAilabsTmallgenieAuthRefreshAPIResponseModel is 刷新token 成功返回结果
type AlibabaAilabsTmallgenieAuthRefreshAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_refresh_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// access token过期时间，相对时间，单位为秒
	AccessExpiresIn int64 `json:"access_expires_in,omitempty" xml:"access_expires_in,omitempty"`
	// access_token
	AccessToken string `json:"access_token,omitempty" xml:"access_token,omitempty"`
	// refresh token过期时间，相对时间，单位为秒
	RefreshExpiresIn int64 `json:"refresh_expires_in,omitempty" xml:"refresh_expires_in,omitempty"`
	// refresh_token
	RefreshToken string `json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
}
