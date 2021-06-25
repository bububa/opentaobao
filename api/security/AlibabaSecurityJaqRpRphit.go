package security

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/security"
)

/* 
聚安全-实人认证日志打点接口 
alibaba.security.jaq.rp.rphit

聚安全实人认证日志打点接口
*/
func AlibabaSecurityJaqRpRphit(clt *core.SDKClient, req *security.AlibabaSecurityJaqRpRphitRequest, session string) (*security.AlibabaSecurityJaqRpRphitResponse, error) {
    var resp security.AlibabaSecurityJaqRpRphitAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
