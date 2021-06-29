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
    clientId   string
    // 授权码
    authCode   string
    // 授权类型，只支持authorization_code
    grantType   string
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
func (r *AlibabaAilabsTmallgenieAuthGettokenRequest) SetClientId(clientId string) error {
    r.clientId = clientId
    r.Set("client_id", clientId)
    return nil
}

// ClientId Getter
func (r AlibabaAilabsTmallgenieAuthGettokenRequest) GetClientId() string {
    return r.clientId
}
// AuthCode Setter
// 授权码
func (r *AlibabaAilabsTmallgenieAuthGettokenRequest) SetAuthCode(authCode string) error {
    r.authCode = authCode
    r.Set("auth_code", authCode)
    return nil
}

// AuthCode Getter
func (r AlibabaAilabsTmallgenieAuthGettokenRequest) GetAuthCode() string {
    return r.authCode
}
// GrantType Setter
// 授权类型，只支持authorization_code
func (r *AlibabaAilabsTmallgenieAuthGettokenRequest) SetGrantType(grantType string) error {
    r.grantType = grantType
    r.Set("grant_type", grantType)
    return nil
}

// GrantType Getter
func (r AlibabaAilabsTmallgenieAuthGettokenRequest) GetGrantType() string {
    return r.grantType
}
