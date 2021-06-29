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
    _uuid   string
    // 设备昵称
    _nickName   string
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
func (r *AlibabaAlinkDeviceInfoUpdateRequest) SetUuid(_uuid string) error {
    r._uuid = _uuid
    r.Set("uuid", _uuid)
    return nil
}

// Uuid Getter
func (r AlibabaAlinkDeviceInfoUpdateRequest) GetUuid() string {
    return r._uuid
}
// NickName Setter
// 设备昵称
func (r *AlibabaAlinkDeviceInfoUpdateRequest) SetNickName(_nickName string) error {
    r._nickName = _nickName
    r.Set("nick_name", _nickName)
    return nil
}

// NickName Getter
func (r AlibabaAlinkDeviceInfoUpdateRequest) GetNickName() string {
    return r._nickName
}
