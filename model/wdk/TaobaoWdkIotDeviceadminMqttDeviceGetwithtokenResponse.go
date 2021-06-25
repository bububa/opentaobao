package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取mqtt设备信息 APIResponse
taobao.wdk.iot.deviceadmin.mqtt.device.getwithtoken

智能硬件设备动态注册和获取mqtt设备信息
*/
type TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenResponse `json:"taobao_wdk_iot_deviceadmin_mqtt_device_getwithtoken_response,omitempty"`
}

type TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenResponse struct {

    // result
    Result   *HmResult `json:"result,omitempty"`

}
