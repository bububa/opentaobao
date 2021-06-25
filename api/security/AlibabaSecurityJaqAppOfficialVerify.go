package security

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/security"
)

/* 
聚安全验证官方应用接口 
alibaba.security.jaq.app.official.verify

接入用户来查询应用是否为官方应用
*/
func AlibabaSecurityJaqAppOfficialVerify(clt *core.SDKClient, req *security.AlibabaSecurityJaqAppOfficialVerifyRequest, session string) (*security.AlibabaSecurityJaqAppOfficialVerifyResponse, error) {
    var resp security.AlibabaSecurityJaqAppOfficialVerifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
