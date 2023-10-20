package alilabs

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIResponseModel).Reset()
}

// AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIResponseModel is 扫码激活设备 成功返回结果
type AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_device_qrcode_activate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果消息
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// 结果对象
	Result *SCanQrCodeResultVo `json:"result,omitempty" xml:"result,omitempty"`
	// 结果码
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RetMsg = ""
	m.Result = nil
	m.RetCode = 0
}

var poolAlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIResponse)
	},
}

// GetAlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIResponse 从 sync.Pool 获取 AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIResponse
func GetAlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIResponse() *AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIResponse {
	return poolAlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIResponse.Get().(*AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIResponse)
}

// ReleaseAlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIResponse 将 AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIResponse(v *AlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIResponse) {
	v.Reset()
	poolAlibabaAilabsTmallgenieAuthDeviceQrcodeActivateAPIResponse.Put(v)
}
