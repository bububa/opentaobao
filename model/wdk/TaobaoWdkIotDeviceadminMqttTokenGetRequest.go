package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取MQTT访问令牌 API请求
taobao.wdk.iot.deviceadmin.mqtt.token.get

智能硬件设备动态注册和获取mqtt设备信息
*/
type TaobaoWdkIotDeviceadminMqttTokenGetRequest struct {
    model.Params
    // accessKey
    _accessKey   string
    // 申请令牌的客户端时间戳
    _applyTimestamp   int64
}

// 初始化TaobaoWdkIotDeviceadminMqttTokenGetRequest对象
func NewTaobaoWdkIotDeviceadminMqttTokenGetRequest() *TaobaoWdkIotDeviceadminMqttTokenGetRequest{
    return &TaobaoWdkIotDeviceadminMqttTokenGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWdkIotDeviceadminMqttTokenGetRequest) GetApiMethodName() string {
    return "taobao.wdk.iot.deviceadmin.mqtt.token.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWdkIotDeviceadminMqttTokenGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AccessKey Setter
// accessKey
func (r *TaobaoWdkIotDeviceadminMqttTokenGetRequest) SetAccessKey(_accessKey string) error {
    r._accessKey = _accessKey
    r.Set("access_key", _accessKey)
    return nil
}

// AccessKey Getter
func (r TaobaoWdkIotDeviceadminMqttTokenGetRequest) GetAccessKey() string {
    return r._accessKey
}
// ApplyTimestamp Setter
// 申请令牌的客户端时间戳
func (r *TaobaoWdkIotDeviceadminMqttTokenGetRequest) SetApplyTimestamp(_applyTimestamp int64) error {
    r._applyTimestamp = _applyTimestamp
    r.Set("apply_timestamp", _applyTimestamp)
    return nil
}

// ApplyTimestamp Getter
func (r TaobaoWdkIotDeviceadminMqttTokenGetRequest) GetApplyTimestamp() int64 {
    return r._applyTimestamp
}
