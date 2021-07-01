package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractLoginAlipayauthAPIResponse
双11到店互动花呗红包获取token鉴权接口 API返回值
alibaba.interact.login.alipayauth

双11到店互动花呗红包获取token鉴权接口 */
type AlibabaInteractLoginAlipayauthAPIResponse struct {
	model.CommonResponse
	AlibabaInteractLoginAlipayauthAPIResponseModel
}

// AlibabaInteractLoginAlipayauthAPIResponseModel is 双11到店互动花呗红包获取token鉴权接口 成功返回结果
type AlibabaInteractLoginAlipayauthAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_login_alipayauth_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
