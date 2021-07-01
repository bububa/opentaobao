package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备详细信息 API请求
taobao.ailab.aicloud.top.device.detailinfo.get

获取设备详细信息
*/
type TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest struct {
    model.Params
    // 三方用户id或淘宝openId
    _originUserId   string
    // 账号秘钥
    _schemaKey   string
    // 三方传extUser，淘宝传openTaoBaoUser
    _userType   string
    // 设备id
    _deviceId   string
}

// 初始化TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest对象
func NewTaobaoAilabAicloudTopDeviceDetailinfoGetRequest() *TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest{
    return &TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.detailinfo.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OriginUserId Setter
// 三方用户id或淘宝openId
func (r *TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest) SetOriginUserId(_originUserId string) error {
    r._originUserId = _originUserId
    r.Set("origin_user_id", _originUserId)
    return nil
}

// OriginUserId Getter
func (r TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest) GetOriginUserId() string {
    return r._originUserId
}
// SchemaKey Setter
// 账号秘钥
func (r *TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest) SetSchemaKey(_schemaKey string) error {
    r._schemaKey = _schemaKey
    r.Set("schema_key", _schemaKey)
    return nil
}

// SchemaKey Getter
func (r TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest) GetSchemaKey() string {
    return r._schemaKey
}
// UserType Setter
// 三方传extUser，淘宝传openTaoBaoUser
func (r *TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest) SetUserType(_userType string) error {
    r._userType = _userType
    r.Set("user_type", _userType)
    return nil
}

// UserType Getter
func (r TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest) GetUserType() string {
    return r._userType
}
// DeviceId Setter
// 设备id
func (r *TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest) SetDeviceId(_deviceId string) error {
    r._deviceId = _deviceId
    r.Set("device_id", _deviceId)
    return nil
}

// DeviceId Getter
func (r TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest) GetDeviceId() string {
    return r._deviceId
}
