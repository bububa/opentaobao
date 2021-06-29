package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备详情 API请求
alibaba.ailabs.tmallgenie.auth.device.get

通过此接口获取设备详情
*/
type AlibabaAilabsTmallgenieAuthDeviceGetRequest struct {
    model.Params
    // 客户id
    _clientId   string
    // 用户开放id
    _userOpenId   string
    // 设备uuid
    _uuid   string
}

// 初始化AlibabaAilabsTmallgenieAuthDeviceGetRequest对象
func NewAlibabaAilabsTmallgenieAuthDeviceGetRequest() *AlibabaAilabsTmallgenieAuthDeviceGetRequest{
    return &AlibabaAilabsTmallgenieAuthDeviceGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthDeviceGetRequest) GetApiMethodName() string {
    return "alibaba.ailabs.tmallgenie.auth.device.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthDeviceGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ClientId Setter
// 客户id
func (r *AlibabaAilabsTmallgenieAuthDeviceGetRequest) SetClientId(_clientId string) error {
    r._clientId = _clientId
    r.Set("client_id", _clientId)
    return nil
}

// ClientId Getter
func (r AlibabaAilabsTmallgenieAuthDeviceGetRequest) GetClientId() string {
    return r._clientId
}
// UserOpenId Setter
// 用户开放id
func (r *AlibabaAilabsTmallgenieAuthDeviceGetRequest) SetUserOpenId(_userOpenId string) error {
    r._userOpenId = _userOpenId
    r.Set("user_open_id", _userOpenId)
    return nil
}

// UserOpenId Getter
func (r AlibabaAilabsTmallgenieAuthDeviceGetRequest) GetUserOpenId() string {
    return r._userOpenId
}
// Uuid Setter
// 设备uuid
func (r *AlibabaAilabsTmallgenieAuthDeviceGetRequest) SetUuid(_uuid string) error {
    r._uuid = _uuid
    r.Set("uuid", _uuid)
    return nil
}

// Uuid Getter
func (r AlibabaAilabsTmallgenieAuthDeviceGetRequest) GetUuid() string {
    return r._uuid
}
