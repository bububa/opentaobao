package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取MQTT访问令牌 API返回值 
taobao.wdk.iot.deviceadmin.mqtt.token.get

智能硬件设备动态注册和获取mqtt设备信息
*/
type TaobaoWdkIotDeviceadminMqttTokenGetAPIResponse struct {
    model.CommonResponse
    TaobaoWdkIotDeviceadminMqttTokenGetAPIResponseModel
}

// 获取MQTT访问令牌 成功返回结果
type TaobaoWdkIotDeviceadminMqttTokenGetAPIResponseModel struct {
    XMLName xml.Name `xml:"wdk_iot_deviceadmin_mqtt_token_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TaobaoWdkIotDeviceadminMqttTokenGetHmResult `json:"result,omitempty" xml:"result,omitempty"`
}
