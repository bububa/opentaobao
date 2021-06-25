package security

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/security"
)

/* 
ocr同时实名校验 
alibaba.security.jaq.rp.cloud.ocr.check

聚安全实人认证证件OCR识别功能接口
*/
func AlibabaSecurityJaqRpCloudOcrCheck(clt *core.SDKClient, req *security.AlibabaSecurityJaqRpCloudOcrCheckRequest, session string) (*security.AlibabaSecurityJaqRpCloudOcrCheckResponse, error) {
    var resp security.AlibabaSecurityJaqRpCloudOcrCheckAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
