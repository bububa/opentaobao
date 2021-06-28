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
    // Response *TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenResponse `json:"wdk_iot_deviceadmin_mqtt_device_getwithtoken_response,omitempty"` 
    TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenResponse
}

/* model for simplify = false
type TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenResponse struct {

    // result
    
    Result  *struct {
        HmResult  *HmResult `json:"hm_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenResponse struct {

    // result
    Result   *HmResult `json:"result,omitempty"`

}
