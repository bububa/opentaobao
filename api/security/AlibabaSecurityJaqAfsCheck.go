package security

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/security"
)

/* 
反欺诈二次验证接口 
alibaba.security.jaq.afs.check

反欺诈二次验证接口
*/
func AlibabaSecurityJaqAfsCheck(clt *core.SDKClient, req *security.AlibabaSecurityJaqAfsCheckRequest, session string) (*security.AlibabaSecurityJaqAfsCheckAPIResponse, error) {
    var resp security.AlibabaSecurityJaqAfsCheckAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
