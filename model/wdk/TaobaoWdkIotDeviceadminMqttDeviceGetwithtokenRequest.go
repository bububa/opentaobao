package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取mqtt设备信息 API请求
taobao.wdk.iot.deviceadmin.mqtt.device.getwithtoken

智能硬件设备动态注册和获取mqtt设备信息
*/
type TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenRequest struct {
    model.Params
    // 业务编码，具体编码请联系杜尘
    businessCode   int64
    // 设备类型编码，具体编码请联系杜尘
    deviceType   int64
    // 环境编码，0日常，1预发，2线上
    enviroCode   int64
    // 设备的唯一标识码，比如网卡的MAC地址
    deviceId   string
    // 访问令牌
    applyAccessToken   string
}

// 初始化TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenRequest对象
func NewTaobaoWdkIotDeviceadminMqttDeviceGetwithtokenRequest() *TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenRequest{
    return &TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenRequest) GetApiMethodName() string {
    return "taobao.wdk.iot.deviceadmin.mqtt.device.getwithtoken"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BusinessCode Setter
// 业务编码，具体编码请联系杜尘
func (r *TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenRequest) SetBusinessCode(businessCode int64) error {
    r.businessCode = businessCode
    r.Set("business_code", businessCode)
    return nil
}

// BusinessCode Getter
func (r TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenRequest) GetBusinessCode() int64 {
    return r.businessCode
}
// DeviceType Setter
// 设备类型编码，具体编码请联系杜尘
func (r *TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenRequest) SetDeviceType(deviceType int64) error {
    r.deviceType = deviceType
    r.Set("device_type", deviceType)
    return nil
}

// DeviceType Getter
func (r TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenRequest) GetDeviceType() int64 {
    return r.deviceType
}
// EnviroCode Setter
// 环境编码，0日常，1预发，2线上
func (r *TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenRequest) SetEnviroCode(enviroCode int64) error {
    r.enviroCode = enviroCode
    r.Set("enviro_code", enviroCode)
    return nil
}

// EnviroCode Getter
func (r TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenRequest) GetEnviroCode() int64 {
    return r.enviroCode
}
// DeviceId Setter
// 设备的唯一标识码，比如网卡的MAC地址
func (r *TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

// DeviceId Getter
func (r TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenRequest) GetDeviceId() string {
    return r.deviceId
}
// ApplyAccessToken Setter
// 访问令牌
func (r *TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenRequest) SetApplyAccessToken(applyAccessToken string) error {
    r.applyAccessToken = applyAccessToken
    r.Set("apply_access_token", applyAccessToken)
    return nil
}

// ApplyAccessToken Getter
func (r TaobaoWdkIotDeviceadminMqttDeviceGetwithtokenRequest) GetApplyAccessToken() string {
    return r.applyAccessToken
}
