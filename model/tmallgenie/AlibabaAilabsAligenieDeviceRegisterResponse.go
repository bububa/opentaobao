package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵开放平台获取设备秘钥 APIResponse
alibaba.ailabs.aligenie.device.register

向天猫精灵inside平台注册设备mac地址，并获取设备的唯一密钥
*/
type AlibabaAilabsAligenieDeviceRegisterAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsAligenieDeviceRegisterResponse
}

type AlibabaAilabsAligenieDeviceRegisterResponse struct {
    XMLName xml.Name `xml:"alibaba_ailabs_aligenie_device_register_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 设备秘钥信息
    
    DeviceSecretInfos   []DeviceSecretInfo `json:"device_secret_infos,omitempty" xml:"device_secret_infos>device_secret_info,omitempty"`
    
    
}
