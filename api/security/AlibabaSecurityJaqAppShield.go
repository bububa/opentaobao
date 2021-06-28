package security

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/security"
)

/* 
应用加固接口 
alibaba.security.jaq.app.shield

提交应用进行应用加固,加固后需通过alibaba.security.jaq.app.shieldresult.get接口查询加固结果
*/
func AlibabaSecurityJaqAppShield(clt *core.SDKClient, req *security.AlibabaSecurityJaqAppShieldRequest, session string) (*security.AlibabaSecurityJaqAppShieldAPIResponse, error) {
    var resp security.AlibabaSecurityJaqAppShieldAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
