package security

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/security"
)

/* 
聚安全实人认证获取结果接口 
alibaba.security.jaq.rp.fetchmaterial

聚安全实人认证获取结果接口
*/
func AlibabaSecurityJaqRpFetchmaterial(clt *core.SDKClient, req *security.AlibabaSecurityJaqRpFetchmaterialAPIRequest, session string) (*security.AlibabaSecurityJaqRpFetchmaterialAPIResponse, error) {
    var resp security.AlibabaSecurityJaqRpFetchmaterialAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
