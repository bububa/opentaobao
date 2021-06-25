package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取MQTT访问令牌 APIRequest
taobao.wdk.iot.deviceadmin.mqtt.token.get

智能硬件设备动态注册和获取mqtt设备信息
*/
type TaobaoWdkIotDeviceadminMqttTokenGetRequest struct {
    model.Params

    // accessKey
    accessKey   string 

    // 申请令牌的客户端时间戳
    applyTimestamp   int64 

}

func NewTaobaoWdkIotDeviceadminMqttTokenGetRequest() *TaobaoWdkIotDeviceadminMqttTokenGetRequest{
    return &TaobaoWdkIotDeviceadminMqttTokenGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWdkIotDeviceadminMqttTokenGetRequest) GetApiMethodName() string {
    return "taobao.wdk.iot.deviceadmin.mqtt.token.get"
}

func (r TaobaoWdkIotDeviceadminMqttTokenGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWdkIotDeviceadminMqttTokenGetRequest) SetAccessKey(accessKey string) error {
    r.accessKey = accessKey
    r.Set("access_key", accessKey)
    return nil
}

func (r TaobaoWdkIotDeviceadminMqttTokenGetRequest) GetAccessKey() string {
    return r.accessKey
}

func (r *TaobaoWdkIotDeviceadminMqttTokenGetRequest) SetApplyTimestamp(applyTimestamp int64) error {
    r.applyTimestamp = applyTimestamp
    r.Set("apply_timestamp", applyTimestamp)
    return nil
}

func (r TaobaoWdkIotDeviceadminMqttTokenGetRequest) GetApplyTimestamp() int64 {
    return r.applyTimestamp
}

