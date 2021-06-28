package security

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/security"
)

/* 
聚安全安全验证检查结果获取接口 
alibaba.security.jaq.captcha.verify.result.fetch

获取二次验证的结果
*/
func AlibabaSecurityJaqCaptchaVerifyResultFetch(clt *core.SDKClient, req *security.AlibabaSecurityJaqCaptchaVerifyResultFetchRequest, session string) (*security.AlibabaSecurityJaqCaptchaVerifyResultFetchAPIResponse, error) {
    var resp security.AlibabaSecurityJaqCaptchaVerifyResultFetchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
