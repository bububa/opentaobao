package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设备授权 APIRequest
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

func NewAlibabaAilabsTmallgenieAuthGettokenRequest() *AlibabaAilabsTmallgenieAuthGettokenRequest{
    return &AlibabaAilabsTmallgenieAuthGettokenRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsTmallgenieAuthGettokenRequest) GetApiMethodName() string {
    return "alibaba.ailabs.tmallgenie.auth.gettoken"
}

func (r AlibabaAilabsTmallgenieAuthGettokenRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsTmallgenieAuthGettokenRequest) SetClientId(clientId string) error {
    r.clientId = clientId
    r.Set("client_id", clientId)
    return nil
}

func (r AlibabaAilabsTmallgenieAuthGettokenRequest) GetClientId() string {
    return r.clientId
}

func (r *AlibabaAilabsTmallgenieAuthGettokenRequest) SetAuthCode(authCode string) error {
    r.authCode = authCode
    r.Set("auth_code", authCode)
    return nil
}

func (r AlibabaAilabsTmallgenieAuthGettokenRequest) GetAuthCode() string {
    return r.authCode
}

func (r *AlibabaAilabsTmallgenieAuthGettokenRequest) SetGrantType(grantType string) error {
    r.grantType = grantType
    r.Set("grant_type", grantType)
    return nil
}

func (r AlibabaAilabsTmallgenieAuthGettokenRequest) GetGrantType() string {
    return r.grantType
}

