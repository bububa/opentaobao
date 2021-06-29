package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备状态信息 API请求
taobao.ailab.aicloud.top.device.statusinfo.get

获取设备状态信息
*/
type TaobaoAilabAicloudTopDeviceStatusinfoGetRequest struct {
    model.Params
    // 三方用户id或淘宝openId
    originUserId   string
    // 账号秘钥
    schemaKey   string
    // 三方传extUser，淘宝传openTaoBaoUser
    userType   string
    // 设备id
    deviceId   string
}

// 初始化TaobaoAilabAicloudTopDeviceStatusinfoGetRequest对象
func NewTaobaoAilabAicloudTopDeviceStatusinfoGetRequest() *TaobaoAilabAicloudTopDeviceStatusinfoGetRequest{
    return &TaobaoAilabAicloudTopDeviceStatusinfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceStatusinfoGetRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.statusinfo.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceStatusinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OriginUserId Setter
// 三方用户id或淘宝openId
func (r *TaobaoAilabAicloudTopDeviceStatusinfoGetRequest) SetOriginUserId(originUserId string) error {
    r.originUserId = originUserId
    r.Set("origin_user_id", originUserId)
    return nil
}

// OriginUserId Getter
func (r TaobaoAilabAicloudTopDeviceStatusinfoGetRequest) GetOriginUserId() string {
    return r.originUserId
}
// SchemaKey Setter
// 账号秘钥
func (r *TaobaoAilabAicloudTopDeviceStatusinfoGetRequest) SetSchemaKey(schemaKey string) error {
    r.schemaKey = schemaKey
    r.Set("schema_key", schemaKey)
    return nil
}

// SchemaKey Getter
func (r TaobaoAilabAicloudTopDeviceStatusinfoGetRequest) GetSchemaKey() string {
    return r.schemaKey
}
// UserType Setter
// 三方传extUser，淘宝传openTaoBaoUser
func (r *TaobaoAilabAicloudTopDeviceStatusinfoGetRequest) SetUserType(userType string) error {
    r.userType = userType
    r.Set("user_type", userType)
    return nil
}

// UserType Getter
func (r TaobaoAilabAicloudTopDeviceStatusinfoGetRequest) GetUserType() string {
    return r.userType
}
// DeviceId Setter
// 设备id
func (r *TaobaoAilabAicloudTopDeviceStatusinfoGetRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

// DeviceId Getter
func (r TaobaoAilabAicloudTopDeviceStatusinfoGetRequest) GetDeviceId() string {
    return r.deviceId
}
