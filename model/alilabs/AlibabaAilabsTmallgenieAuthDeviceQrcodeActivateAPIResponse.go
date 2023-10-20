package alilabs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIResponse 扫码激活设备 API返回值
// alibaba.ailabs.tmallgenie.auth.device.qrcode.activate
//
// 三方带屏设备显示二维码（从天猫精灵云获取），使用三方APP扫码，将扫码到的安全code，通过TOP接口请求天猫精灵云，精灵云解析安全code的数据并激活对应的设备。
type AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIResponseModel
}

// AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIResponseModel is 扫码激活设备 成功返回结果
type AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_device_qrcode_activate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果消息
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// 结果对象
	Result *ScanQrCodeResultVo `json:"result,omitempty" xml:"result,omitempty"`
	// 结果码
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
}
