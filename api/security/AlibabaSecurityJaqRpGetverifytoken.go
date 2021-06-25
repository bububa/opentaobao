package security

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/security"
)

/* 
聚安全实人认证获取认证会话token 
alibaba.security.jaq.rp.getverifytoken

聚安全实人认证获取认证会话token
*/
func AlibabaSecurityJaqRpGetverifytoken(clt *core.SDKClient, req *security.AlibabaSecurityJaqRpGetverifytokenRequest, session string) (*security.AlibabaSecurityJaqRpGetverifytokenResponse, error) {
    var resp security.AlibabaSecurityJaqRpGetverifytokenAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
