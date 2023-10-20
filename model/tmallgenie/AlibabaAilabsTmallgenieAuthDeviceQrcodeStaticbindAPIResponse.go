package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIResponse 静态二维码绑定 API返回值
// alibaba.ailabs.tmallgenie.auth.device.qrcode.staticbind
//
// 静态二维码绑定
type AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIResponseModel
}

// AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIResponseModel is 静态二维码绑定 成功返回结果
type AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_device_qrcode_staticbind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindResult `json:"result,omitempty" xml:"result,omitempty"`
}
