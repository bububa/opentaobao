package security

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/security"
)

/* 
聚安全实人认证查询认证结果 
alibaba.security.jaq.rp.query

聚安全实人认证查询认证结果
*/
func AlibabaSecurityJaqRpQuery(clt *core.SDKClient, req *security.AlibabaSecurityJaqRpQueryRequest, session string) (*security.AlibabaSecurityJaqRpQueryResponse, error) {
    var resp security.AlibabaSecurityJaqRpQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
