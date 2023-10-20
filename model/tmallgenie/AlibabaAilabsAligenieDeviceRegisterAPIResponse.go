package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsAligenieDeviceRegisterAPIResponse 天猫精灵开放平台获取设备秘钥 API返回值
// alibaba.ailabs.aligenie.device.register
//
// 向天猫精灵inside平台注册设备mac地址，并获取设备的唯一密钥
type AlibabaAilabsAligenieDeviceRegisterAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsAligenieDeviceRegisterAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsAligenieDeviceRegisterAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsAligenieDeviceRegisterAPIResponseModel).Reset()
}

// AlibabaAilabsAligenieDeviceRegisterAPIResponseModel is 天猫精灵开放平台获取设备秘钥 成功返回结果
type AlibabaAilabsAligenieDeviceRegisterAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_aligenie_device_register_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 设备秘钥信息
	DeviceSecretInfos []DeviceSecretInfo `json:"device_secret_infos,omitempty" xml:"device_secret_infos>device_secret_info,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsAligenieDeviceRegisterAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DeviceSecretInfos = m.DeviceSecretInfos[:0]
}

var poolAlibabaAilabsAligenieDeviceRegisterAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsAligenieDeviceRegisterAPIResponse)
	},
}

// GetAlibabaAilabsAligenieDeviceRegisterAPIResponse 从 sync.Pool 获取 AlibabaAilabsAligenieDeviceRegisterAPIResponse
func GetAlibabaAilabsAligenieDeviceRegisterAPIResponse() *AlibabaAilabsAligenieDeviceRegisterAPIResponse {
	return poolAlibabaAilabsAligenieDeviceRegisterAPIResponse.Get().(*AlibabaAilabsAligenieDeviceRegisterAPIResponse)
}

// ReleaseAlibabaAilabsAligenieDeviceRegisterAPIResponse 将 AlibabaAilabsAligenieDeviceRegisterAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsAligenieDeviceRegisterAPIResponse(v *AlibabaAilabsAligenieDeviceRegisterAPIResponse) {
	v.Reset()
	poolAlibabaAilabsAligenieDeviceRegisterAPIResponse.Put(v)
}
