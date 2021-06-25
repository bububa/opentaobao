package alink

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
绑定设备 APIRequest
alibaba.alink.device.bind

阿里智能解绑设备
*/
type AlibabaAlinkDeviceBindRequest struct {
    model.Params

    // 设备id
    uuid   string 

}

func NewAlibabaAlinkDeviceBindRequest() *AlibabaAlinkDeviceBindRequest{
    return &AlibabaAlinkDeviceBindRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlinkDeviceBindRequest) GetApiMethodName() string {
    return "alibaba.alink.device.bind"
}

func (r AlibabaAlinkDeviceBindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlinkDeviceBindRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r AlibabaAlinkDeviceBindRequest) GetUuid() string {
    return r.uuid
}

