package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备详情 APIRequest
alibaba.ailabs.tmallgenie.auth.device.get

通过此接口获取设备详情
*/
type AlibabaAilabsTmallgenieAuthDeviceGetRequest struct {
    model.Params

    // 客户id
    clientId   string 

    // 用户开放id
    userOpenId   string 

    // 设备uuid
    uuid   string 

}

func NewAlibabaAilabsTmallgenieAuthDeviceGetRequest() *AlibabaAilabsTmallgenieAuthDeviceGetRequest{
    return &AlibabaAilabsTmallgenieAuthDeviceGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsTmallgenieAuthDeviceGetRequest) GetApiMethodName() string {
    return "alibaba.ailabs.tmallgenie.auth.device.get"
}

func (r AlibabaAilabsTmallgenieAuthDeviceGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsTmallgenieAuthDeviceGetRequest) SetClientId(clientId string) error {
    r.clientId = clientId
    r.Set("client_id", clientId)
    return nil
}

func (r AlibabaAilabsTmallgenieAuthDeviceGetRequest) GetClientId() string {
    return r.clientId
}

func (r *AlibabaAilabsTmallgenieAuthDeviceGetRequest) SetUserOpenId(userOpenId string) error {
    r.userOpenId = userOpenId
    r.Set("user_open_id", userOpenId)
    return nil
}

func (r AlibabaAilabsTmallgenieAuthDeviceGetRequest) GetUserOpenId() string {
    return r.userOpenId
}

func (r *AlibabaAilabsTmallgenieAuthDeviceGetRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r AlibabaAilabsTmallgenieAuthDeviceGetRequest) GetUuid() string {
    return r.uuid
}

