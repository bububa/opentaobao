package security

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/security"
)

/* 
聚安全实人认证提交认证接口 
alibaba.security.jaq.rp.submit

聚安全实人认证提交认证接口
*/
func AlibabaSecurityJaqRpSubmit(clt *core.SDKClient, req *security.AlibabaSecurityJaqRpSubmitRequest, session string) (*security.AlibabaSecurityJaqRpSubmitResponse, error) {
    var resp security.AlibabaSecurityJaqRpSubmitAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
