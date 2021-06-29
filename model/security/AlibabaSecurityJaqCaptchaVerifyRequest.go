package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全安全验证检查接口 API请求
alibaba.security.jaq.captcha.verify

聚安全安全验证检查
*/
type AlibabaSecurityJaqCaptchaVerifyRequest struct {
    model.Params
    // 验证码发送渠道类型 1-短信 2-语音 3-邮件
    _captchaType   int64
    // 扩展字段，格式为JSON字符串，用于传递“短信验证”等验证方式所需的额外入参，例如用户输入的验证码等，格式及JSON字段key定义请参考示例
    _extendData   string
    // 协议版本号
    _protocolVersion   string
    // 软token签名
    _rsign   string
    // 软token索引
    _rtkenIndex   string
    // 验证码会话ID
    _sessionId   string
    // UMID token
    _utoken   string
    // 风险识别接口返回的信息口令
    _infoToken   string
}

// 初始化AlibabaSecurityJaqCaptchaVerifyRequest对象
func NewAlibabaSecurityJaqCaptchaVerifyRequest() *AlibabaSecurityJaqCaptchaVerifyRequest{
    return &AlibabaSecurityJaqCaptchaVerifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqCaptchaVerifyRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.captcha.verify"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqCaptchaVerifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CaptchaType Setter
// 验证码发送渠道类型 1-短信 2-语音 3-邮件
func (r *AlibabaSecurityJaqCaptchaVerifyRequest) SetCaptchaType(_captchaType int64) error {
    r._captchaType = _captchaType
    r.Set("captcha_type", _captchaType)
    return nil
}

// CaptchaType Getter
func (r AlibabaSecurityJaqCaptchaVerifyRequest) GetCaptchaType() int64 {
    return r._captchaType
}
// ExtendData Setter
// 扩展字段，格式为JSON字符串，用于传递“短信验证”等验证方式所需的额外入参，例如用户输入的验证码等，格式及JSON字段key定义请参考示例
func (r *AlibabaSecurityJaqCaptchaVerifyRequest) SetExtendData(_extendData string) error {
    r._extendData = _extendData
    r.Set("extend_data", _extendData)
    return nil
}

// ExtendData Getter
func (r AlibabaSecurityJaqCaptchaVerifyRequest) GetExtendData() string {
    return r._extendData
}
// ProtocolVersion Setter
// 协议版本号
func (r *AlibabaSecurityJaqCaptchaVerifyRequest) SetProtocolVersion(_protocolVersion string) error {
    r._protocolVersion = _protocolVersion
    r.Set("protocol_version", _protocolVersion)
    return nil
}

// ProtocolVersion Getter
func (r AlibabaSecurityJaqCaptchaVerifyRequest) GetProtocolVersion() string {
    return r._protocolVersion
}
// Rsign Setter
// 软token签名
func (r *AlibabaSecurityJaqCaptchaVerifyRequest) SetRsign(_rsign string) error {
    r._rsign = _rsign
    r.Set("rsign", _rsign)
    return nil
}

// Rsign Getter
func (r AlibabaSecurityJaqCaptchaVerifyRequest) GetRsign() string {
    return r._rsign
}
// RtkenIndex Setter
// 软token索引
func (r *AlibabaSecurityJaqCaptchaVerifyRequest) SetRtkenIndex(_rtkenIndex string) error {
    r._rtkenIndex = _rtkenIndex
    r.Set("rtken_index", _rtkenIndex)
    return nil
}

// RtkenIndex Getter
func (r AlibabaSecurityJaqCaptchaVerifyRequest) GetRtkenIndex() string {
    return r._rtkenIndex
}
// SessionId Setter
// 验证码会话ID
func (r *AlibabaSecurityJaqCaptchaVerifyRequest) SetSessionId(_sessionId string) error {
    r._sessionId = _sessionId
    r.Set("session_id", _sessionId)
    return nil
}

// SessionId Getter
func (r AlibabaSecurityJaqCaptchaVerifyRequest) GetSessionId() string {
    return r._sessionId
}
// Utoken Setter
// UMID token
func (r *AlibabaSecurityJaqCaptchaVerifyRequest) SetUtoken(_utoken string) error {
    r._utoken = _utoken
    r.Set("utoken", _utoken)
    return nil
}

// Utoken Getter
func (r AlibabaSecurityJaqCaptchaVerifyRequest) GetUtoken() string {
    return r._utoken
}
// InfoToken Setter
// 风险识别接口返回的信息口令
func (r *AlibabaSecurityJaqCaptchaVerifyRequest) SetInfoToken(_infoToken string) error {
    r._infoToken = _infoToken
    r.Set("info_token", _infoToken)
    return nil
}

// InfoToken Getter
func (r AlibabaSecurityJaqCaptchaVerifyRequest) GetInfoToken() string {
    return r._infoToken
}
