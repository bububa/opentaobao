package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设备授权 API请求
alibaba.ailabs.tmallgenie.auth.gettoken

获取设备授权码
*/
type AlibabaAilabsTmallgenieAuthGettokenRequest struct {
    model.Params
    // clientId
    _clientId   string
    // 授权码
    _authCode   string
    // 授权类型，只支持authorization_code
    _grantType   string
}

// 初始化AlibabaAilabsTmallgenieAuthGettokenRequest对象
func NewAlibabaAilabsTmallgenieAuthGettokenRequest() *AlibabaAilabsTmallgenieAuthGettokenRequest{
    return &AlibabaAilabsTmallgenieAuthGettokenRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthGettokenRequest) GetApiMethodName() string {
    return "alibaba.ailabs.tmallgenie.auth.gettoken"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthGettokenRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ClientId Setter
// clientId
func (r *AlibabaAilabsTmallgenieAuthGettokenRequest) SetClientId(_clientId string) error {
    r._clientId = _clientId
    r.Set("client_id", _clientId)
    return nil
}

// ClientId Getter
func (r AlibabaAilabsTmallgenieAuthGettokenRequest) GetClientId() string {
    return r._clientId
}
// AuthCode Setter
// 授权码
func (r *AlibabaAilabsTmallgenieAuthGettokenRequest) SetAuthCode(_authCode string) error {
    r._authCode = _authCode
    r.Set("auth_code", _authCode)
    return nil
}

// AuthCode Getter
func (r AlibabaAilabsTmallgenieAuthGettokenRequest) GetAuthCode() string {
    return r._authCode
}
// GrantType Setter
// 授权类型，只支持authorization_code
func (r *AlibabaAilabsTmallgenieAuthGettokenRequest) SetGrantType(_grantType string) error {
    r._grantType = _grantType
    r.Set("grant_type", _grantType)
    return nil
}

// GrantType Getter
func (r AlibabaAilabsTmallgenieAuthGettokenRequest) GetGrantType() string {
    return r._grantType
}
