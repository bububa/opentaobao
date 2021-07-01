package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备扩展信息 API请求
taobao.ailab.aicloud.top.device.extinfo.get

获取设备扩展信息
*/
type TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest struct {
    model.Params
    // 三方id、淘宝openId
    _originUserId   string
    // 账号秘钥
    _schemaKey   string
    // 类型：openTaoBao, extUser
    _userType   string
    // 设备id
    _deviceId   string
}

// 初始化TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest对象
func NewTaobaoAilabAicloudTopDeviceExtinfoGetRequest() *TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest{
    return &TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.extinfo.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OriginUserId Setter
// 三方id、淘宝openId
func (r *TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest) SetOriginUserId(_originUserId string) error {
    r._originUserId = _originUserId
    r.Set("origin_user_id", _originUserId)
    return nil
}

// OriginUserId Getter
func (r TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest) GetOriginUserId() string {
    return r._originUserId
}
// SchemaKey Setter
// 账号秘钥
func (r *TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest) SetSchemaKey(_schemaKey string) error {
    r._schemaKey = _schemaKey
    r.Set("schema_key", _schemaKey)
    return nil
}

// SchemaKey Getter
func (r TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest) GetSchemaKey() string {
    return r._schemaKey
}
// UserType Setter
// 类型：openTaoBao, extUser
func (r *TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest) SetUserType(_userType string) error {
    r._userType = _userType
    r.Set("user_type", _userType)
    return nil
}

// UserType Getter
func (r TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest) GetUserType() string {
    return r._userType
}
// DeviceId Setter
// 设备id
func (r *TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest) SetDeviceId(_deviceId string) error {
    r._deviceId = _deviceId
    r.Set("device_id", _deviceId)
    return nil
}

// DeviceId Getter
func (r TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest) GetDeviceId() string {
    return r._deviceId
}
