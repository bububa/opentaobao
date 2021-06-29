package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
解绑设备 API请求
alibaba.ailabs.tmallgenie.auth.device.unbind

通过此接口解绑天猫精灵设备
*/
type AlibabaAilabsTmallgenieAuthDeviceUnbindRequest struct {
    model.Params
    // 客户id
    _clientId   string
    // 用户开放id
    _userOpenId   string
    // 设备uuid
    _uuid   string
}

// 初始化AlibabaAilabsTmallgenieAuthDeviceUnbindRequest对象
func NewAlibabaAilabsTmallgenieAuthDeviceUnbindRequest() *AlibabaAilabsTmallgenieAuthDeviceUnbindRequest{
    return &AlibabaAilabsTmallgenieAuthDeviceUnbindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthDeviceUnbindRequest) GetApiMethodName() string {
    return "alibaba.ailabs.tmallgenie.auth.device.unbind"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthDeviceUnbindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ClientId Setter
// 客户id
func (r *AlibabaAilabsTmallgenieAuthDeviceUnbindRequest) SetClientId(_clientId string) error {
    r._clientId = _clientId
    r.Set("client_id", _clientId)
    return nil
}

// ClientId Getter
func (r AlibabaAilabsTmallgenieAuthDeviceUnbindRequest) GetClientId() string {
    return r._clientId
}
// UserOpenId Setter
// 用户开放id
func (r *AlibabaAilabsTmallgenieAuthDeviceUnbindRequest) SetUserOpenId(_userOpenId string) error {
    r._userOpenId = _userOpenId
    r.Set("user_open_id", _userOpenId)
    return nil
}

// UserOpenId Getter
func (r AlibabaAilabsTmallgenieAuthDeviceUnbindRequest) GetUserOpenId() string {
    return r._userOpenId
}
// Uuid Setter
// 设备uuid
func (r *AlibabaAilabsTmallgenieAuthDeviceUnbindRequest) SetUuid(_uuid string) error {
    r._uuid = _uuid
    r.Set("uuid", _uuid)
    return nil
}

// Uuid Getter
func (r AlibabaAilabsTmallgenieAuthDeviceUnbindRequest) GetUuid() string {
    return r._uuid
}
