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
type TaobaoAilabAicloudTopDeviceDetailinfoGetRequest struct {
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

// 初始化TaobaoAilabAicloudTopDeviceDetailinfoGetRequest对象
func NewTaobaoAilabAicloudTopDeviceDetailinfoGetRequest() *TaobaoAilabAicloudTopDeviceDetailinfoGetRequest{
    return &TaobaoAilabAicloudTopDeviceDetailinfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceDetailinfoGetRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.detailinfo.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceDetailinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OriginUserId Setter
// 三方用户id或淘宝openId
func (r *TaobaoAilabAicloudTopDeviceDetailinfoGetRequest) SetOriginUserId(_originUserId string) error {
    r._originUserId = _originUserId
    r.Set("origin_user_id", _originUserId)
    return nil
}

// OriginUserId Getter
func (r TaobaoAilabAicloudTopDeviceDetailinfoGetRequest) GetOriginUserId() string {
    return r._originUserId
}
// SchemaKey Setter
// 账号秘钥
func (r *TaobaoAilabAicloudTopDeviceDetailinfoGetRequest) SetSchemaKey(_schemaKey string) error {
    r._schemaKey = _schemaKey
    r.Set("schema_key", _schemaKey)
    return nil
}

// SchemaKey Getter
func (r TaobaoAilabAicloudTopDeviceDetailinfoGetRequest) GetSchemaKey() string {
    return r._schemaKey
}
// UserType Setter
// 三方传extUser，淘宝传openTaoBaoUser
func (r *TaobaoAilabAicloudTopDeviceDetailinfoGetRequest) SetUserType(_userType string) error {
    r._userType = _userType
    r.Set("user_type", _userType)
    return nil
}

// UserType Getter
func (r TaobaoAilabAicloudTopDeviceDetailinfoGetRequest) GetUserType() string {
    return r._userType
}
// DeviceId Setter
// 设备id
func (r *TaobaoAilabAicloudTopDeviceDetailinfoGetRequest) SetDeviceId(_deviceId string) error {
    r._deviceId = _deviceId
    r.Set("device_id", _deviceId)
    return nil
}

// DeviceId Getter
func (r TaobaoAilabAicloudTopDeviceDetailinfoGetRequest) GetDeviceId() string {
    return r._deviceId
}
