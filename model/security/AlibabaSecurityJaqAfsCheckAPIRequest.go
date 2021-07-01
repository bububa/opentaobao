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
type AlibabaSecurityJaqAfsCheckAPIRequest struct {
    model.Params
    // 上报平台枚举值 1标识Android端 2标识iOS端 3标识PC端及其他
    _platform   int64
    // token，来自客户端上报
    _token   string
    // 会话ID，来自客户端上报
    _sessionId   string
    // 签名串，来自客户端上报
    _sig   string
    // 用户接入的时候获取的风控key
    _afsKey   string
}

// 初始化AlibabaSecurityJaqAfsCheckAPIRequest对象
func NewAlibabaSecurityJaqAfsCheckRequest() *AlibabaSecurityJaqAfsCheckAPIRequest{
    return &AlibabaSecurityJaqAfsCheckAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqAfsCheckAPIRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.afs.check"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqAfsCheckAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Platform Setter
// 上报平台枚举值 1标识Android端 2标识iOS端 3标识PC端及其他
func (r *AlibabaSecurityJaqAfsCheckAPIRequest) SetPlatform(_platform int64) error {
    r._platform = _platform
    r.Set("platform", _platform)
    return nil
}

// Platform Getter
func (r AlibabaSecurityJaqAfsCheckAPIRequest) GetPlatform() int64 {
    return r._platform
}
// Token Setter
// token，来自客户端上报
func (r *AlibabaSecurityJaqAfsCheckAPIRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r AlibabaSecurityJaqAfsCheckAPIRequest) GetToken() string {
    return r._token
}
// SessionId Setter
// 会话ID，来自客户端上报
func (r *AlibabaSecurityJaqAfsCheckAPIRequest) SetSessionId(_sessionId string) error {
    r._sessionId = _sessionId
    r.Set("session_id", _sessionId)
    return nil
}

// SessionId Getter
func (r AlibabaSecurityJaqAfsCheckAPIRequest) GetSessionId() string {
    return r._sessionId
}
// Sig Setter
// 签名串，来自客户端上报
func (r *AlibabaSecurityJaqAfsCheckAPIRequest) SetSig(_sig string) error {
    r._sig = _sig
    r.Set("sig", _sig)
    return nil
}

// Sig Getter
func (r AlibabaSecurityJaqAfsCheckAPIRequest) GetSig() string {
    return r._sig
}
// AfsKey Setter
// 用户接入的时候获取的风控key
func (r *AlibabaSecurityJaqAfsCheckAPIRequest) SetAfsKey(_afsKey string) error {
    r._afsKey = _afsKey
    r.Set("afs_key", _afsKey)
    return nil
}

// AfsKey Getter
func (r AlibabaSecurityJaqAfsCheckAPIRequest) GetAfsKey() string {
    return r._afsKey
}
