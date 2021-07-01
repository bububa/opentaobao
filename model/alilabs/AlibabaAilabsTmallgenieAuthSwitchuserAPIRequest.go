package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
切换用户 API请求
alibaba.ailabs.tmallgenie.auth.switchuser

设备切换授权用户
*/
type AlibabaAilabsTmallgenieAuthSwitchuserAPIRequest struct {
    model.Params
    // client_id
    _clientId   string
    // 目标用户openId
    _newUserOpenId   string
    // 当前拥有设备权限的用户openId
    _oldUserOpenId   string
    // 设备uuid
    _uuid   string
}

// 初始化AlibabaAilabsTmallgenieAuthSwitchuserAPIRequest对象
func NewAlibabaAilabsTmallgenieAuthSwitchuserRequest() *AlibabaAilabsTmallgenieAuthSwitchuserAPIRequest{
    return &AlibabaAilabsTmallgenieAuthSwitchuserAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthSwitchuserAPIRequest) GetApiMethodName() string {
    return "alibaba.ailabs.tmallgenie.auth.switchuser"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthSwitchuserAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ClientId Setter
// client_id
func (r *AlibabaAilabsTmallgenieAuthSwitchuserAPIRequest) SetClientId(_clientId string) error {
    r._clientId = _clientId
    r.Set("client_id", _clientId)
    return nil
}

// ClientId Getter
func (r AlibabaAilabsTmallgenieAuthSwitchuserAPIRequest) GetClientId() string {
    return r._clientId
}
// NewUserOpenId Setter
// 目标用户openId
func (r *AlibabaAilabsTmallgenieAuthSwitchuserAPIRequest) SetNewUserOpenId(_newUserOpenId string) error {
    r._newUserOpenId = _newUserOpenId
    r.Set("new_user_open_id", _newUserOpenId)
    return nil
}

// NewUserOpenId Getter
func (r AlibabaAilabsTmallgenieAuthSwitchuserAPIRequest) GetNewUserOpenId() string {
    return r._newUserOpenId
}
// OldUserOpenId Setter
// 当前拥有设备权限的用户openId
func (r *AlibabaAilabsTmallgenieAuthSwitchuserAPIRequest) SetOldUserOpenId(_oldUserOpenId string) error {
    r._oldUserOpenId = _oldUserOpenId
    r.Set("old_user_open_id", _oldUserOpenId)
    return nil
}

// OldUserOpenId Getter
func (r AlibabaAilabsTmallgenieAuthSwitchuserAPIRequest) GetOldUserOpenId() string {
    return r._oldUserOpenId
}
// Uuid Setter
// 设备uuid
func (r *AlibabaAilabsTmallgenieAuthSwitchuserAPIRequest) SetUuid(_uuid string) error {
    r._uuid = _uuid
    r.Set("uuid", _uuid)
    return nil
}

// Uuid Getter
func (r AlibabaAilabsTmallgenieAuthSwitchuserAPIRequest) GetUuid() string {
    return r._uuid
}
