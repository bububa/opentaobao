package alink

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
解绑设备 APIRequest
alibaba.alink.device.unbind

阿里智能解绑设备
*/
type AlibabaAlinkDeviceUnbindRequest struct {
    model.Params

    // 设备id
    uuid   string 

}

func NewAlibabaAlinkDeviceUnbindRequest() *AlibabaAlinkDeviceUnbindRequest{
    return &AlibabaAlinkDeviceUnbindRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlinkDeviceUnbindRequest) GetApiMethodName() string {
    return "alibaba.alink.device.unbind"
}

func (r AlibabaAlinkDeviceUnbindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlinkDeviceUnbindRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r AlibabaAlinkDeviceUnbindRequest) GetUuid() string {
    return r.uuid
}

