package alink

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新设备昵称等信息 APIRequest
alibaba.alink.device.info.update

更新设备昵称等信息
*/
type AlibabaAlinkDeviceInfoUpdateRequest struct {
    model.Params

    // 设备id
    uuid   string 

    // 设备昵称
    nickName   string 

}

func NewAlibabaAlinkDeviceInfoUpdateRequest() *AlibabaAlinkDeviceInfoUpdateRequest{
    return &AlibabaAlinkDeviceInfoUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlinkDeviceInfoUpdateRequest) GetApiMethodName() string {
    return "alibaba.alink.device.info.update"
}

func (r AlibabaAlinkDeviceInfoUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlinkDeviceInfoUpdateRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r AlibabaAlinkDeviceInfoUpdateRequest) GetUuid() string {
    return r.uuid
}

func (r *AlibabaAlinkDeviceInfoUpdateRequest) SetNickName(nickName string) error {
    r.nickName = nickName
    r.Set("nick_name", nickName)
    return nil
}

func (r AlibabaAlinkDeviceInfoUpdateRequest) GetNickName() string {
    return r.nickName
}

