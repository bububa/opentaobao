package security

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/security"
)

/* 
聚安全实人认证开始 
alibaba.security.jaq.rp.start

聚安全实人认证开始
*/
func AlibabaSecurityJaqRpStart(clt *core.SDKClient, req *security.AlibabaSecurityJaqRpStartRequest, session string) (*security.AlibabaSecurityJaqRpStartResponse, error) {
    var resp security.AlibabaSecurityJaqRpStartAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
