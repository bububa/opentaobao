package security

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/security"
)

/* 
聚安全图文识别同步检测接口 
alibaba.security.jaq.ocr.image.sync.detect

图像字符识别同步检测接口
*/
func AlibabaSecurityJaqOcrImageSyncDetect(clt *core.SDKClient, req *security.AlibabaSecurityJaqOcrImageSyncDetectRequest, session string) (*security.AlibabaSecurityJaqOcrImageSyncDetectResponse, error) {
    var resp security.AlibabaSecurityJaqOcrImageSyncDetectAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
