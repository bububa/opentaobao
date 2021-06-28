package security

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/security"
)

/* 
聚安全实人认证证件OCR识别功能接口 
alibaba.security.jaq.rp.ocr

聚安全实人认证证件OCR识别功能接口
*/
func AlibabaSecurityJaqRpOcr(clt *core.SDKClient, req *security.AlibabaSecurityJaqRpOcrRequest, session string) (*security.AlibabaSecurityJaqRpOcrAPIResponse, error) {
    var resp security.AlibabaSecurityJaqRpOcrAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
