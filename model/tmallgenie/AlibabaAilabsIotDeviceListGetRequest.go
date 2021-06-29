package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取iot设备列表 APIRequest
alibaba.ailabs.iot.device.list.get

通过此接口获取用户名下的iot设备列表
*/
type AlibabaAilabsIotDeviceListGetRequest struct {
    model.Params

    // 用户open id
    userOpenId   string 

    // client id
    clientId   string 

}

func NewAlibabaAilabsIotDeviceListGetRequest() *AlibabaAilabsIotDeviceListGetRequest{
    return &AlibabaAilabsIotDeviceListGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsIotDeviceListGetRequest) GetApiMethodName() string {
    return "alibaba.ailabs.iot.device.list.get"
}

func (r AlibabaAilabsIotDeviceListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsIotDeviceListGetRequest) SetUserOpenId(userOpenId string) error {
    r.userOpenId = userOpenId
    r.Set("user_open_id", userOpenId)
    return nil
}

func (r AlibabaAilabsIotDeviceListGetRequest) GetUserOpenId() string {
    return r.userOpenId
}

func (r *AlibabaAilabsIotDeviceListGetRequest) SetClientId(clientId string) error {
    r.clientId = clientId
    r.Set("client_id", clientId)
    return nil
}

func (r AlibabaAilabsIotDeviceListGetRequest) GetClientId() string {
    return r.clientId
}

