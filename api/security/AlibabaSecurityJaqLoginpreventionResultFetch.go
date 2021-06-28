package security

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/security"
)

/* 
获取登录保护结果 
alibaba.security.jaq.loginprevention.result.fetch

获取登录保护结果
*/
func AlibabaSecurityJaqLoginpreventionResultFetch(clt *core.SDKClient, req *security.AlibabaSecurityJaqLoginpreventionResultFetchRequest, session string) (*security.AlibabaSecurityJaqLoginpreventionResultFetchAPIResponse, error) {
    var resp security.AlibabaSecurityJaqLoginpreventionResultFetchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
