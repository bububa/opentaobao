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
type TaobaoAilabAicloudTopDeviceExtinfoGetRequest struct {
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

// 初始化TaobaoAilabAicloudTopDeviceExtinfoGetRequest对象
func NewTaobaoAilabAicloudTopDeviceExtinfoGetRequest() *TaobaoAilabAicloudTopDeviceExtinfoGetRequest{
    return &TaobaoAilabAicloudTopDeviceExtinfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceExtinfoGetRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.extinfo.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceExtinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OriginUserId Setter
// 三方id、淘宝openId
func (r *TaobaoAilabAicloudTopDeviceExtinfoGetRequest) SetOriginUserId(_originUserId string) error {
    r._originUserId = _originUserId
    r.Set("origin_user_id", _originUserId)
    return nil
}

// OriginUserId Getter
func (r TaobaoAilabAicloudTopDeviceExtinfoGetRequest) GetOriginUserId() string {
    return r._originUserId
}
// SchemaKey Setter
// 账号秘钥
func (r *TaobaoAilabAicloudTopDeviceExtinfoGetRequest) SetSchemaKey(_schemaKey string) error {
    r._schemaKey = _schemaKey
    r.Set("schema_key", _schemaKey)
    return nil
}

// SchemaKey Getter
func (r TaobaoAilabAicloudTopDeviceExtinfoGetRequest) GetSchemaKey() string {
    return r._schemaKey
}
// UserType Setter
// 类型：openTaoBao, extUser
func (r *TaobaoAilabAicloudTopDeviceExtinfoGetRequest) SetUserType(_userType string) error {
    r._userType = _userType
    r.Set("user_type", _userType)
    return nil
}

// UserType Getter
func (r TaobaoAilabAicloudTopDeviceExtinfoGetRequest) GetUserType() string {
    return r._userType
}
// DeviceId Setter
// 设备id
func (r *TaobaoAilabAicloudTopDeviceExtinfoGetRequest) SetDeviceId(_deviceId string) error {
    r._deviceId = _deviceId
    r.Set("device_id", _deviceId)
    return nil
}

// DeviceId Getter
func (r TaobaoAilabAicloudTopDeviceExtinfoGetRequest) GetDeviceId() string {
    return r._deviceId
}
