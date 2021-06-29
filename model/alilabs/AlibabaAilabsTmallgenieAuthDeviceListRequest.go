package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户设备列表 APIRequest
alibaba.ailabs.tmallgenie.auth.device.list

通过此接口获取用户绑定的设备信息列表
*/
type AlibabaAilabsTmallgenieAuthDeviceListRequest struct {
    model.Params

    // 客户id
    clientId   string 

    // 用户开放id
    userOpenId   string 

}

func NewAlibabaAilabsTmallgenieAuthDeviceListRequest() *AlibabaAilabsTmallgenieAuthDeviceListRequest{
    return &AlibabaAilabsTmallgenieAuthDeviceListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsTmallgenieAuthDeviceListRequest) GetApiMethodName() string {
    return "alibaba.ailabs.tmallgenie.auth.device.list"
}

func (r AlibabaAilabsTmallgenieAuthDeviceListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsTmallgenieAuthDeviceListRequest) SetClientId(clientId string) error {
    r.clientId = clientId
    r.Set("client_id", clientId)
    return nil
}

func (r AlibabaAilabsTmallgenieAuthDeviceListRequest) GetClientId() string {
    return r.clientId
}

func (r *AlibabaAilabsTmallgenieAuthDeviceListRequest) SetUserOpenId(userOpenId string) error {
    r.userOpenId = userOpenId
    r.Set("user_open_id", userOpenId)
    return nil
}

func (r AlibabaAilabsTmallgenieAuthDeviceListRequest) GetUserOpenId() string {
    return r.userOpenId
}

