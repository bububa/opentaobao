package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全安全验证检查结果获取接口 API请求
alibaba.security.jaq.captcha.verify.result.fetch

获取二次验证的结果
*/
type AlibabaSecurityJaqCaptchaVerifyResultFetchRequest struct {
    model.Params
    // 二次验证获取验证检查结果所需的seesionId
    _sessionId   string
}

// 初始化AlibabaSecurityJaqCaptchaVerifyResultFetchRequest对象
func NewAlibabaSecurityJaqCaptchaVerifyResultFetchRequest() *AlibabaSecurityJaqCaptchaVerifyResultFetchRequest{
    return &AlibabaSecurityJaqCaptchaVerifyResultFetchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqCaptchaVerifyResultFetchRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.captcha.verify.result.fetch"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqCaptchaVerifyResultFetchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SessionId Setter
// 二次验证获取验证检查结果所需的seesionId
func (r *AlibabaSecurityJaqCaptchaVerifyResultFetchRequest) SetSessionId(_sessionId string) error {
    r._sessionId = _sessionId
    r.Set("session_id", _sessionId)
    return nil
}

// SessionId Getter
func (r AlibabaSecurityJaqCaptchaVerifyResultFetchRequest) GetSessionId() string {
    return r._sessionId
}
