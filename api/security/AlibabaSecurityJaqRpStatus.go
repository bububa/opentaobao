package security

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/security"
)

/* 
聚安全实人认证查询状态接口 
alibaba.security.jaq.rp.status

聚安全实人认证查询状态接口
*/
func AlibabaSecurityJaqRpStatus(clt *core.SDKClient, req *security.AlibabaSecurityJaqRpStatusRequest, session string) (*security.AlibabaSecurityJaqRpStatusResponse, error) {
    var resp security.AlibabaSecurityJaqRpStatusAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
