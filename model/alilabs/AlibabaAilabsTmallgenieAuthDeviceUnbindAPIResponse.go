package alilabs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsTmallgenieAuthDeviceUnbindAPIResponse
解绑设备 API返回值
alibaba.ailabs.tmallgenie.auth.device.unbind

通过此接口解绑天猫精灵设备 */
type AlibabaAilabsTmallgenieAuthDeviceUnbindAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsTmallgenieAuthDeviceUnbindAPIResponseModel
}

// AlibabaAilabsTmallgenieAuthDeviceUnbindAPIResponseModel is 解绑设备 成功返回结果
type AlibabaAilabsTmallgenieAuthDeviceUnbindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_device_unbind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回包装类
	Result *AlibabaAilabsTmallgenieAuthDeviceUnbindResult `json:"result,omitempty" xml:"result,omitempty"`
}
