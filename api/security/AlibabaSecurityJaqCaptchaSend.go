package security

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/security"
)

/* 
聚安全安全验证发起接口 
alibaba.security.jaq.captcha.send

聚安全安全验证发起
*/
func AlibabaSecurityJaqCaptchaSend(clt *core.SDKClient, req *security.AlibabaSecurityJaqCaptchaSendRequest, session string) (*security.AlibabaSecurityJaqCaptchaSendResponse, error) {
    var resp security.AlibabaSecurityJaqCaptchaSendAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
