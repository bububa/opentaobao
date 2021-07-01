package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

/* AlibabaSecurityJaqUrlScan
恶意网址检测接口
alibaba.security.jaq.url.scan

url扫描接口 */
func AlibabaSecurityJaqUrlScan(clt *core.SDKClient, req *security.AlibabaSecurityJaqUrlScanAPIRequest, session string) (*security.AlibabaSecurityJaqUrlScanAPIResponse, error) {
	var resp security.AlibabaSecurityJaqUrlScanAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
