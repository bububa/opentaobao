package security

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/security"
)

/* 
验证姓名和证件号 
alibaba.security.jaq.rp.cloud.realname.check

验证姓名和证件号
*/
func AlibabaSecurityJaqRpCloudRealnameCheck(clt *core.SDKClient, req *security.AlibabaSecurityJaqRpCloudRealnameCheckAPIRequest, session string) (*security.AlibabaSecurityJaqRpCloudRealnameCheckAPIResponse, error) {
    var resp security.AlibabaSecurityJaqRpCloudRealnameCheckAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
