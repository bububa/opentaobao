package alink

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新设备昵称等信息 API请求
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

// 初始化AlibabaAlinkDeviceInfoUpdateRequest对象
func NewAlibabaAlinkDeviceInfoUpdateRequest() *AlibabaAlinkDeviceInfoUpdateRequest{
    return &AlibabaAlinkDeviceInfoUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlinkDeviceInfoUpdateRequest) GetApiMethodName() string {
    return "alibaba.alink.device.info.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlinkDeviceInfoUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Uuid Setter
// 设备id
func (r *AlibabaAlinkDeviceInfoUpdateRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

// Uuid Getter
func (r AlibabaAlinkDeviceInfoUpdateRequest) GetUuid() string {
    return r.uuid
}
// NickName Setter
// 设备昵称
func (r *AlibabaAlinkDeviceInfoUpdateRequest) SetNickName(nickName string) error {
    r.nickName = nickName
    r.Set("nick_name", nickName)
    return nil
}

// NickName Getter
func (r AlibabaAlinkDeviceInfoUpdateRequest) GetNickName() string {
    return r.nickName
}
