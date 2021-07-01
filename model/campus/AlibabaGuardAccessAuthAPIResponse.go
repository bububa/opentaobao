package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaGuardAccessAuthAPIResponse
鉴权 API返回值
alibaba.guard.access.auth

刷卡鉴权 */
type AlibabaGuardAccessAuthAPIResponse struct {
	model.CommonResponse
	AlibabaGuardAccessAuthAPIResponseModel
}

// AlibabaGuardAccessAuthAPIResponseModel is 鉴权 成功返回结果
type AlibabaGuardAccessAuthAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_guard_access_auth_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
