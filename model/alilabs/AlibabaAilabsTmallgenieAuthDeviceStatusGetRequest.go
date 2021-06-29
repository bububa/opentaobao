package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设备状态查询 APIRequest
alibaba.ailabs.tmallgenie.auth.device.status.get

提供给天猫精灵定制机厂商 查询设备在线状态值
*/
type AlibabaAilabsTmallgenieAuthDeviceStatusGetRequest struct {
    model.Params

    // 给产品分配的唯一ID（22位）
    clientId   string 

    // 精灵用户的唯一外部ID
    userOpenId   string 

    // 精灵设备的唯一ID
    uuid   string 

}

func NewAlibabaAilabsTmallgenieAuthDeviceStatusGetRequest() *AlibabaAilabsTmallgenieAuthDeviceStatusGetRequest{
    return &AlibabaAilabsTmallgenieAuthDeviceStatusGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsTmallgenieAuthDeviceStatusGetRequest) GetApiMethodName() string {
    return "alibaba.ailabs.tmallgenie.auth.device.status.get"
}

func (r AlibabaAilabsTmallgenieAuthDeviceStatusGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsTmallgenieAuthDeviceStatusGetRequest) SetClientId(clientId string) error {
    r.clientId = clientId
    r.Set("client_id", clientId)
    return nil
}

func (r AlibabaAilabsTmallgenieAuthDeviceStatusGetRequest) GetClientId() string {
    return r.clientId
}

func (r *AlibabaAilabsTmallgenieAuthDeviceStatusGetRequest) SetUserOpenId(userOpenId string) error {
    r.userOpenId = userOpenId
    r.Set("user_open_id", userOpenId)
    return nil
}

func (r AlibabaAilabsTmallgenieAuthDeviceStatusGetRequest) GetUserOpenId() string {
    return r.userOpenId
}

func (r *AlibabaAilabsTmallgenieAuthDeviceStatusGetRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r AlibabaAilabsTmallgenieAuthDeviceStatusGetRequest) GetUuid() string {
    return r.uuid
}

