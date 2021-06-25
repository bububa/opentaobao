package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全安全验证发起接口 APIRequest
alibaba.security.jaq.captcha.send

聚安全安全验证发起
*/
type AlibabaSecurityJaqCaptchaSendRequest struct {
    model.Params

    // 验证码发送渠道类型 1-短信 3-邮件
    captchaType   int64 

    // 扩展字段，格式为JSON字符串，用于传递“滚小球”等验证方式所需的额外入参，例如屏幕尺寸等，请参考示例
    extendData   string 

    // 协议版本号
    protocolVersion   string 

    // 软token签名
    rsign   string 

    // 软token索引
    rtkenIndex   string 

    // UMID token
    utoken   string 

    // 风险识别接口返回的信息口令
    infoToken   string 

}

func NewAlibabaSecurityJaqCaptchaSendRequest() *AlibabaSecurityJaqCaptchaSendRequest{
    return &AlibabaSecurityJaqCaptchaSendRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqCaptchaSendRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.captcha.send"
}

func (r AlibabaSecurityJaqCaptchaSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqCaptchaSendRequest) SetCaptchaType(captchaType int64) error {
    r.captchaType = captchaType
    r.Set("captcha_type", captchaType)
    return nil
}

func (r AlibabaSecurityJaqCaptchaSendRequest) GetCaptchaType() int64 {
    return r.captchaType
}

func (r *AlibabaSecurityJaqCaptchaSendRequest) SetExtendData(extendData string) error {
    r.extendData = extendData
    r.Set("extend_data", extendData)
    return nil
}

func (r AlibabaSecurityJaqCaptchaSendRequest) GetExtendData() string {
    return r.extendData
}

func (r *AlibabaSecurityJaqCaptchaSendRequest) SetProtocolVersion(protocolVersion string) error {
    r.protocolVersion = protocolVersion
    r.Set("protocol_version", protocolVersion)
    return nil
}

func (r AlibabaSecurityJaqCaptchaSendRequest) GetProtocolVersion() string {
    return r.protocolVersion
}

func (r *AlibabaSecurityJaqCaptchaSendRequest) SetRsign(rsign string) error {
    r.rsign = rsign
    r.Set("rsign", rsign)
    return nil
}

func (r AlibabaSecurityJaqCaptchaSendRequest) GetRsign() string {
    return r.rsign
}

func (r *AlibabaSecurityJaqCaptchaSendRequest) SetRtkenIndex(rtkenIndex string) error {
    r.rtkenIndex = rtkenIndex
    r.Set("rtken_index", rtkenIndex)
    return nil
}

func (r AlibabaSecurityJaqCaptchaSendRequest) GetRtkenIndex() string {
    return r.rtkenIndex
}

func (r *AlibabaSecurityJaqCaptchaSendRequest) SetUtoken(utoken string) error {
    r.utoken = utoken
    r.Set("utoken", utoken)
    return nil
}

func (r AlibabaSecurityJaqCaptchaSendRequest) GetUtoken() string {
    return r.utoken
}

func (r *AlibabaSecurityJaqCaptchaSendRequest) SetInfoToken(infoToken string) error {
    r.infoToken = infoToken
    r.Set("info_token", infoToken)
    return nil
}

func (r AlibabaSecurityJaqCaptchaSendRequest) GetInfoToken() string {
    return r.infoToken
}

