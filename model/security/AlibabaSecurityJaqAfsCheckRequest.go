package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
反欺诈二次验证接口 API请求
alibaba.security.jaq.afs.check

反欺诈二次验证接口
*/
type AlibabaSecurityJaqAfsCheckRequest struct {
    model.Params
    // 上报平台枚举值 1标识Android端 2标识iOS端 3标识PC端及其他
    platform   int64
    // token，来自客户端上报
    token   string
    // 会话ID，来自客户端上报
    sessionId   string
    // 签名串，来自客户端上报
    sig   string
    // 用户接入的时候获取的风控key
    afsKey   string
}

// 初始化AlibabaSecurityJaqAfsCheckRequest对象
func NewAlibabaSecurityJaqAfsCheckRequest() *AlibabaSecurityJaqAfsCheckRequest{
    return &AlibabaSecurityJaqAfsCheckRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqAfsCheckRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.afs.check"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqAfsCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Platform Setter
// 上报平台枚举值 1标识Android端 2标识iOS端 3标识PC端及其他
func (r *AlibabaSecurityJaqAfsCheckRequest) SetPlatform(platform int64) error {
    r.platform = platform
    r.Set("platform", platform)
    return nil
}

// Platform Getter
func (r AlibabaSecurityJaqAfsCheckRequest) GetPlatform() int64 {
    return r.platform
}
// Token Setter
// token，来自客户端上报
func (r *AlibabaSecurityJaqAfsCheckRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r AlibabaSecurityJaqAfsCheckRequest) GetToken() string {
    return r.token
}
// SessionId Setter
// 会话ID，来自客户端上报
func (r *AlibabaSecurityJaqAfsCheckRequest) SetSessionId(sessionId string) error {
    r.sessionId = sessionId
    r.Set("session_id", sessionId)
    return nil
}

// SessionId Getter
func (r AlibabaSecurityJaqAfsCheckRequest) GetSessionId() string {
    return r.sessionId
}
// Sig Setter
// 签名串，来自客户端上报
func (r *AlibabaSecurityJaqAfsCheckRequest) SetSig(sig string) error {
    r.sig = sig
    r.Set("sig", sig)
    return nil
}

// Sig Getter
func (r AlibabaSecurityJaqAfsCheckRequest) GetSig() string {
    return r.sig
}
// AfsKey Setter
// 用户接入的时候获取的风控key
func (r *AlibabaSecurityJaqAfsCheckRequest) SetAfsKey(afsKey string) error {
    r.afsKey = afsKey
    r.Set("afs_key", afsKey)
    return nil
}

// AfsKey Getter
func (r AlibabaSecurityJaqAfsCheckRequest) GetAfsKey() string {
    return r.afsKey
}
