package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全安全验证发起接口 API请求
alibaba.security.jaq.captcha.send

聚安全安全验证发起
*/
type AlibabaSecurityJaqCaptchaSendRequest struct {
    model.Params
    // 验证码发送渠道类型 1-短信 3-邮件
    _captchaType   int64
    // 扩展字段，格式为JSON字符串，用于传递“滚小球”等验证方式所需的额外入参，例如屏幕尺寸等，请参考示例
    _extendData   string
    // 协议版本号
    _protocolVersion   string
    // 软token签名
    _rsign   string
    // 软token索引
    _rtkenIndex   string
    // UMID token
    _utoken   string
    // 风险识别接口返回的信息口令
    _infoToken   string
}

// 初始化AlibabaSecurityJaqCaptchaSendRequest对象
func NewAlibabaSecurityJaqCaptchaSendRequest() *AlibabaSecurityJaqCaptchaSendRequest{
    return &AlibabaSecurityJaqCaptchaSendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqCaptchaSendRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.captcha.send"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqCaptchaSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CaptchaType Setter
// 验证码发送渠道类型 1-短信 3-邮件
func (r *AlibabaSecurityJaqCaptchaSendRequest) SetCaptchaType(_captchaType int64) error {
    r._captchaType = _captchaType
    r.Set("captcha_type", _captchaType)
    return nil
}

// CaptchaType Getter
func (r AlibabaSecurityJaqCaptchaSendRequest) GetCaptchaType() int64 {
    return r._captchaType
}
// ExtendData Setter
// 扩展字段，格式为JSON字符串，用于传递“滚小球”等验证方式所需的额外入参，例如屏幕尺寸等，请参考示例
func (r *AlibabaSecurityJaqCaptchaSendRequest) SetExtendData(_extendData string) error {
    r._extendData = _extendData
    r.Set("extend_data", _extendData)
    return nil
}

// ExtendData Getter
func (r AlibabaSecurityJaqCaptchaSendRequest) GetExtendData() string {
    return r._extendData
}
// ProtocolVersion Setter
// 协议版本号
func (r *AlibabaSecurityJaqCaptchaSendRequest) SetProtocolVersion(_protocolVersion string) error {
    r._protocolVersion = _protocolVersion
    r.Set("protocol_version", _protocolVersion)
    return nil
}

// ProtocolVersion Getter
func (r AlibabaSecurityJaqCaptchaSendRequest) GetProtocolVersion() string {
    return r._protocolVersion
}
// Rsign Setter
// 软token签名
func (r *AlibabaSecurityJaqCaptchaSendRequest) SetRsign(_rsign string) error {
    r._rsign = _rsign
    r.Set("rsign", _rsign)
    return nil
}

// Rsign Getter
func (r AlibabaSecurityJaqCaptchaSendRequest) GetRsign() string {
    return r._rsign
}
// RtkenIndex Setter
// 软token索引
func (r *AlibabaSecurityJaqCaptchaSendRequest) SetRtkenIndex(_rtkenIndex string) error {
    r._rtkenIndex = _rtkenIndex
    r.Set("rtken_index", _rtkenIndex)
    return nil
}

// RtkenIndex Getter
func (r AlibabaSecurityJaqCaptchaSendRequest) GetRtkenIndex() string {
    return r._rtkenIndex
}
// Utoken Setter
// UMID token
func (r *AlibabaSecurityJaqCaptchaSendRequest) SetUtoken(_utoken string) error {
    r._utoken = _utoken
    r.Set("utoken", _utoken)
    return nil
}

// Utoken Getter
func (r AlibabaSecurityJaqCaptchaSendRequest) GetUtoken() string {
    return r._utoken
}
// InfoToken Setter
// 风险识别接口返回的信息口令
func (r *AlibabaSecurityJaqCaptchaSendRequest) SetInfoToken(_infoToken string) error {
    r._infoToken = _infoToken
    r.Set("info_token", _infoToken)
    return nil
}

// InfoToken Getter
func (r AlibabaSecurityJaqCaptchaSendRequest) GetInfoToken() string {
    return r._infoToken
}
