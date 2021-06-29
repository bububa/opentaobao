package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
解绑设备 APIRequest
alibaba.ailabs.tmallgenie.auth.device.unbind

通过此接口解绑天猫精灵设备
*/
type AlibabaAilabsTmallgenieAuthDeviceUnbindRequest struct {
    model.Params

    // 客户id
    clientId   string 

    // 用户开放id
    userOpenId   string 

    // 设备uuid
    uuid   string 

}

func NewAlibabaAilabsTmallgenieAuthDeviceUnbindRequest() *AlibabaAilabsTmallgenieAuthDeviceUnbindRequest{
    return &AlibabaAilabsTmallgenieAuthDeviceUnbindRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsTmallgenieAuthDeviceUnbindRequest) GetApiMethodName() string {
    return "alibaba.ailabs.tmallgenie.auth.device.unbind"
}

func (r AlibabaAilabsTmallgenieAuthDeviceUnbindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsTmallgenieAuthDeviceUnbindRequest) SetClientId(clientId string) error {
    r.clientId = clientId
    r.Set("client_id", clientId)
    return nil
}

func (r AlibabaAilabsTmallgenieAuthDeviceUnbindRequest) GetClientId() string {
    return r.clientId
}

func (r *AlibabaAilabsTmallgenieAuthDeviceUnbindRequest) SetUserOpenId(userOpenId string) error {
    r.userOpenId = userOpenId
    r.Set("user_open_id", userOpenId)
    return nil
}

func (r AlibabaAilabsTmallgenieAuthDeviceUnbindRequest) GetUserOpenId() string {
    return r.userOpenId
}

func (r *AlibabaAilabsTmallgenieAuthDeviceUnbindRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r AlibabaAilabsTmallgenieAuthDeviceUnbindRequest) GetUuid() string {
    return r.uuid
}

