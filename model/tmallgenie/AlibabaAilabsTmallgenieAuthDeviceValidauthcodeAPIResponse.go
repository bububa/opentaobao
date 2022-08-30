package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieAuthDeviceValidauthcodeAPIResponse 根据authcode查询绑定结果 API返回值
// alibaba.ailabs.tmallgenie.auth.device.validauthcode
//
// 根据authcode查询绑定结果
type AlibabaAilabsTmallgenieAuthDeviceValidauthcodeAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsTmallgenieAuthDeviceValidauthcodeAPIResponseModel
}

// AlibabaAilabsTmallgenieAuthDeviceValidauthcodeAPIResponseModel is 根据authcode查询绑定结果 成功返回结果
type AlibabaAilabsTmallgenieAuthDeviceValidauthcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_device_validauthcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAilabsTmallgenieAuthDeviceValidauthcodeResult `json:"result,omitempty" xml:"result,omitempty"`
}
