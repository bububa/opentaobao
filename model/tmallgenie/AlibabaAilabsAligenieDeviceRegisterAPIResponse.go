package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabsaligeniedeviceregisterAPIResponse 天猫精灵开放平台获取设备秘钥 API返回值
// alibaba.ailabs.aligenie.device.register
//
// 向天猫精灵inside平台注册设备mac地址，并获取设备的唯一密钥
type AlibabaailabsaligeniedeviceregisterAPIResponse struct {
	model.CommonResponse
	AlibabaailabsaligeniedeviceregisterAPIResponseModel
}

// AlibabaailabsaligeniedeviceregisterAPIResponseModel is 天猫精灵开放平台获取设备秘钥 成功返回结果
type AlibabaailabsaligeniedeviceregisterAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_aligenie_device_register_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 设备秘钥信息
	DeviceSecretInfos []DeviceSecretInfo `json:"device_secret_infos,omitempty" xml:"device_secret_infos>device_secret_info,omitempty"`
}
