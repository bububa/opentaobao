package security

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/security"
)

/* 
获取虚假注册保护结果 
alibaba.security.jaq.spamregisterprevention.result.fetch.new

获取虚假注册保护结果
*/
func AlibabaSecurityJaqSpamregisterpreventionResultFetchNew(clt *core.SDKClient, req *security.AlibabaSecurityJaqSpamregisterpreventionResultFetchNewRequest, session string) (*security.AlibabaSecurityJaqSpamregisterpreventionResultFetchNewResponse, error) {
    var resp security.AlibabaSecurityJaqSpamregisterpreventionResultFetchNewAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
