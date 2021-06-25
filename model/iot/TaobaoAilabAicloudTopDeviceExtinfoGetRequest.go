package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备扩展信息 APIRequest
taobao.ailab.aicloud.top.device.extinfo.get

获取设备扩展信息
*/
type TaobaoAilabAicloudTopDeviceExtinfoGetRequest struct {
    model.Params

    // 三方id、淘宝openId
    originUserId   string 

    // 账号秘钥
    schemaKey   string 

    // 类型：openTaoBao, extUser
    userType   string 

    // 设备id
    deviceId   string 

}

func NewTaobaoAilabAicloudTopDeviceExtinfoGetRequest() *TaobaoAilabAicloudTopDeviceExtinfoGetRequest{
    return &TaobaoAilabAicloudTopDeviceExtinfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopDeviceExtinfoGetRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.extinfo.get"
}

func (r TaobaoAilabAicloudTopDeviceExtinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopDeviceExtinfoGetRequest) SetOriginUserId(originUserId string) error {
    r.originUserId = originUserId
    r.Set("origin_user_id", originUserId)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceExtinfoGetRequest) GetOriginUserId() string {
    return r.originUserId
}

func (r *TaobaoAilabAicloudTopDeviceExtinfoGetRequest) SetSchemaKey(schemaKey string) error {
    r.schemaKey = schemaKey
    r.Set("schema_key", schemaKey)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceExtinfoGetRequest) GetSchemaKey() string {
    return r.schemaKey
}

func (r *TaobaoAilabAicloudTopDeviceExtinfoGetRequest) SetUserType(userType string) error {
    r.userType = userType
    r.Set("user_type", userType)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceExtinfoGetRequest) GetUserType() string {
    return r.userType
}

func (r *TaobaoAilabAicloudTopDeviceExtinfoGetRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceExtinfoGetRequest) GetDeviceId() string {
    return r.deviceId
}

