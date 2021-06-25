package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全安全验证检查结果获取接口 APIRequest
alibaba.security.jaq.captcha.verify.result.fetch

获取二次验证的结果
*/
type AlibabaSecurityJaqCaptchaVerifyResultFetchRequest struct {
    model.Params

    // 二次验证获取验证检查结果所需的seesionId
    sessionId   string 

}

func NewAlibabaSecurityJaqCaptchaVerifyResultFetchRequest() *AlibabaSecurityJaqCaptchaVerifyResultFetchRequest{
    return &AlibabaSecurityJaqCaptchaVerifyResultFetchRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqCaptchaVerifyResultFetchRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.captcha.verify.result.fetch"
}

func (r AlibabaSecurityJaqCaptchaVerifyResultFetchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqCaptchaVerifyResultFetchRequest) SetSessionId(sessionId string) error {
    r.sessionId = sessionId
    r.Set("session_id", sessionId)
    return nil
}

func (r AlibabaSecurityJaqCaptchaVerifyResultFetchRequest) GetSessionId() string {
    return r.sessionId
}

