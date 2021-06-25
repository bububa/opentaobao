package alink

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备详情 APIRequest
alibaba.alink.device.detail.get

阿里智能获取设备详情
*/
type AlibabaAlinkDeviceDetailGetRequest struct {
    model.Params

    // 设备id
    uuid   string 

}

func NewAlibabaAlinkDeviceDetailGetRequest() *AlibabaAlinkDeviceDetailGetRequest{
    return &AlibabaAlinkDeviceDetailGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlinkDeviceDetailGetRequest) GetApiMethodName() string {
    return "alibaba.alink.device.detail.get"
}

func (r AlibabaAlinkDeviceDetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlinkDeviceDetailGetRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r AlibabaAlinkDeviceDetailGetRequest) GetUuid() string {
    return r.uuid
}

