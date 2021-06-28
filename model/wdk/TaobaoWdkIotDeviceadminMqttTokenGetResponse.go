package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取MQTT访问令牌 APIResponse
taobao.wdk.iot.deviceadmin.mqtt.token.get

智能硬件设备动态注册和获取mqtt设备信息
*/
type TaobaoWdkIotDeviceadminMqttTokenGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWdkIotDeviceadminMqttTokenGetResponse `json:"wdk_iot_deviceadmin_mqtt_token_get_response,omitempty"` 
    TaobaoWdkIotDeviceadminMqttTokenGetResponse
}

/* model for simplify = false
type TaobaoWdkIotDeviceadminMqttTokenGetResponse struct {

    // result
    
    Result  *struct {
        HmResult  *HmResult `json:"hm_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoWdkIotDeviceadminMqttTokenGetResponse struct {

    // result
    Result   *HmResult `json:"result,omitempty"`

}
