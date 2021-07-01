package security

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/security"
)

/* 
用户查询加固结果 
alibaba.security.jaq.app.shieldresult.get

用户通过alibaba.security.jaq.app.shield接口提交应用加固后,通过该接口查询加固结果,下载加固包
*/
func AlibabaSecurityJaqAppShieldresultGet(clt *core.SDKClient, req *security.AlibabaSecurityJaqAppShieldresultGetAPIRequest, session string) (*security.AlibabaSecurityJaqAppShieldresultGetAPIResponse, error) {
    var resp security.AlibabaSecurityJaqAppShieldresultGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
