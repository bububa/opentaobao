package security

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/security"
)

/* 
恶意网址检测接口 
alibaba.security.jaq.url.scan

url扫描接口
*/
func AlibabaSecurityJaqUrlScan(clt *core.SDKClient, req *security.AlibabaSecurityJaqUrlScanRequest, session string) (*security.AlibabaSecurityJaqUrlScanResponse, error) {
    var resp security.AlibabaSecurityJaqUrlScanAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
