package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备状态信息 APIRequest
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

func NewTaobaoAilabAicloudTopDeviceStatusinfoGetRequest() *TaobaoAilabAicloudTopDeviceStatusinfoGetRequest{
    return &TaobaoAilabAicloudTopDeviceStatusinfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopDeviceStatusinfoGetRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.statusinfo.get"
}

func (r TaobaoAilabAicloudTopDeviceStatusinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopDeviceStatusinfoGetRequest) SetOriginUserId(originUserId string) error {
    r.originUserId = originUserId
    r.Set("origin_user_id", originUserId)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceStatusinfoGetRequest) GetOriginUserId() string {
    return r.originUserId
}

func (r *TaobaoAilabAicloudTopDeviceStatusinfoGetRequest) SetSchemaKey(schemaKey string) error {
    r.schemaKey = schemaKey
    r.Set("schema_key", schemaKey)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceStatusinfoGetRequest) GetSchemaKey() string {
    return r.schemaKey
}

func (r *TaobaoAilabAicloudTopDeviceStatusinfoGetRequest) SetUserType(userType string) error {
    r.userType = userType
    r.Set("user_type", userType)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceStatusinfoGetRequest) GetUserType() string {
    return r.userType
}

func (r *TaobaoAilabAicloudTopDeviceStatusinfoGetRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceStatusinfoGetRequest) GetDeviceId() string {
    return r.deviceId
}

