package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备详细信息 APIRequest
taobao.ailab.aicloud.top.device.detailinfo.get

获取设备详细信息
*/
type TaobaoAilabAicloudTopDeviceDetailinfoGetRequest struct {
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

func NewTaobaoAilabAicloudTopDeviceDetailinfoGetRequest() *TaobaoAilabAicloudTopDeviceDetailinfoGetRequest{
    return &TaobaoAilabAicloudTopDeviceDetailinfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopDeviceDetailinfoGetRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.detailinfo.get"
}

func (r TaobaoAilabAicloudTopDeviceDetailinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopDeviceDetailinfoGetRequest) SetOriginUserId(originUserId string) error {
    r.originUserId = originUserId
    r.Set("origin_user_id", originUserId)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceDetailinfoGetRequest) GetOriginUserId() string {
    return r.originUserId
}

func (r *TaobaoAilabAicloudTopDeviceDetailinfoGetRequest) SetSchemaKey(schemaKey string) error {
    r.schemaKey = schemaKey
    r.Set("schema_key", schemaKey)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceDetailinfoGetRequest) GetSchemaKey() string {
    return r.schemaKey
}

func (r *TaobaoAilabAicloudTopDeviceDetailinfoGetRequest) SetUserType(userType string) error {
    r.userType = userType
    r.Set("user_type", userType)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceDetailinfoGetRequest) GetUserType() string {
    return r.userType
}

func (r *TaobaoAilabAicloudTopDeviceDetailinfoGetRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceDetailinfoGetRequest) GetDeviceId() string {
    return r.deviceId
}

