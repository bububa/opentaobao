package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据三方ID查询设备注册激活信息 APIRequest
alibaba.ailabs.tmallgenie.auth.device.withdeviceid.get

根据三方ID查询设备注册激活信息
*/
type AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetRequest struct {
    model.Params

    // 设备产品ID
    clientId   string 

    // mac地址
    mac   string 

}

func NewAlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetRequest() *AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetRequest{
    return &AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetRequest) GetApiMethodName() string {
    return "alibaba.ailabs.tmallgenie.auth.device.withdeviceid.get"
}

func (r AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetRequest) SetClientId(clientId string) error {
    r.clientId = clientId
    r.Set("client_id", clientId)
    return nil
}

func (r AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetRequest) GetClientId() string {
    return r.clientId
}

func (r *AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetRequest) SetMac(mac string) error {
    r.mac = mac
    r.Set("mac", mac)
    return nil
}

func (r AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetRequest) GetMac() string {
    return r.mac
}

