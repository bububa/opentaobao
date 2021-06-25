package security

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/security"
)

/* 
获取垃圾注册防控结果 
alibaba.security.jaq.spamregisterprevention.result.fetch

获取垃圾注册防控结果
*/
func AlibabaSecurityJaqSpamregisterpreventionResultFetch(clt *core.SDKClient, req *security.AlibabaSecurityJaqSpamregisterpreventionResultFetchRequest, session string) (*security.AlibabaSecurityJaqSpamregisterpreventionResultFetchResponse, error) {
    var resp security.AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
