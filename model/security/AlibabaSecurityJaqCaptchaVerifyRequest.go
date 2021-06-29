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
    captchaType   int64
    // 扩展字段，格式为JSON字符串，用于传递“短信验证”等验证方式所需的额外入参，例如用户输入的验证码等，格式及JSON字段key定义请参考示例
    extendData   string
    // 协议版本号
    protocolVersion   string
    // 软token签名
    rsign   string
    // 软token索引
    rtkenIndex   string
    // 验证码会话ID
    sessionId   string
    // UMID token
    utoken   string
    // 风险识别接口返回的信息口令
    infoToken   string
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
func (r *AlibabaSecurityJaqCaptchaVerifyRequest) SetCaptchaType(captchaType int64) error {
    r.captchaType = captchaType
    r.Set("captcha_type", captchaType)
    return nil
}

// CaptchaType Getter
func (r AlibabaSecurityJaqCaptchaVerifyRequest) GetCaptchaType() int64 {
    return r.captchaType
}
// ExtendData Setter
// 扩展字段，格式为JSON字符串，用于传递“短信验证”等验证方式所需的额外入参，例如用户输入的验证码等，格式及JSON字段key定义请参考示例
func (r *AlibabaSecurityJaqCaptchaVerifyRequest) SetExtendData(extendData string) error {
    r.extendData = extendData
    r.Set("extend_data", extendData)
    return nil
}

// ExtendData Getter
func (r AlibabaSecurityJaqCaptchaVerifyRequest) GetExtendData() string {
    return r.extendData
}
// ProtocolVersion Setter
// 协议版本号
func (r *AlibabaSecurityJaqCaptchaVerifyRequest) SetProtocolVersion(protocolVersion string) error {
    r.protocolVersion = protocolVersion
    r.Set("protocol_version", protocolVersion)
    return nil
}

// ProtocolVersion Getter
func (r AlibabaSecurityJaqCaptchaVerifyRequest) GetProtocolVersion() string {
    return r.protocolVersion
}
// Rsign Setter
// 软token签名
func (r *AlibabaSecurityJaqCaptchaVerifyRequest) SetRsign(rsign string) error {
    r.rsign = rsign
    r.Set("rsign", rsign)
    return nil
}

// Rsign Getter
func (r AlibabaSecurityJaqCaptchaVerifyRequest) GetRsign() string {
    return r.rsign
}
// RtkenIndex Setter
// 软token索引
func (r *AlibabaSecurityJaqCaptchaVerifyRequest) SetRtkenIndex(rtkenIndex string) error {
    r.rtkenIndex = rtkenIndex
    r.Set("rtken_index", rtkenIndex)
    return nil
}

// RtkenIndex Getter
func (r AlibabaSecurityJaqCaptchaVerifyRequest) GetRtkenIndex() string {
    return r.rtkenIndex
}
// SessionId Setter
// 验证码会话ID
func (r *AlibabaSecurityJaqCaptchaVerifyRequest) SetSessionId(sessionId string) error {
    r.sessionId = sessionId
    r.Set("session_id", sessionId)
    return nil
}

// SessionId Getter
func (r AlibabaSecurityJaqCaptchaVerifyRequest) GetSessionId() string {
    return r.sessionId
}
// Utoken Setter
// UMID token
func (r *AlibabaSecurityJaqCaptchaVerifyRequest) SetUtoken(utoken string) error {
    r.utoken = utoken
    r.Set("utoken", utoken)
    return nil
}

// Utoken Getter
func (r AlibabaSecurityJaqCaptchaVerifyRequest) GetUtoken() string {
    return r.utoken
}
// InfoToken Setter
// 风险识别接口返回的信息口令
func (r *AlibabaSecurityJaqCaptchaVerifyRequest) SetInfoToken(infoToken string) error {
    r.infoToken = infoToken
    r.Set("info_token", infoToken)
    return nil
}

// InfoToken Getter
func (r AlibabaSecurityJaqCaptchaVerifyRequest) GetInfoToken() string {
    return r.infoToken
}
