package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户设备列表 API请求
alibaba.ailabs.tmallgenie.auth.device.list

通过此接口获取用户绑定的设备信息列表
*/
type AlibabaAilabsTmallgenieAuthDeviceListRequest struct {
    model.Params
    // 客户id
    _clientId   string
    // 用户开放id
    _userOpenId   string
}

// 初始化AlibabaAilabsTmallgenieAuthDeviceListRequest对象
func NewAlibabaAilabsTmallgenieAuthDeviceListRequest() *AlibabaAilabsTmallgenieAuthDeviceListRequest{
    return &AlibabaAilabsTmallgenieAuthDeviceListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthDeviceListRequest) GetApiMethodName() string {
    return "alibaba.ailabs.tmallgenie.auth.device.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthDeviceListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ClientId Setter
// 客户id
func (r *AlibabaAilabsTmallgenieAuthDeviceListRequest) SetClientId(_clientId string) error {
    r._clientId = _clientId
    r.Set("client_id", _clientId)
    return nil
}

// ClientId Getter
func (r AlibabaAilabsTmallgenieAuthDeviceListRequest) GetClientId() string {
    return r._clientId
}
// UserOpenId Setter
// 用户开放id
func (r *AlibabaAilabsTmallgenieAuthDeviceListRequest) SetUserOpenId(_userOpenId string) error {
    r._userOpenId = _userOpenId
    r.Set("user_open_id", _userOpenId)
    return nil
}

// UserOpenId Getter
func (r AlibabaAilabsTmallgenieAuthDeviceListRequest) GetUserOpenId() string {
    return r._userOpenId
}
