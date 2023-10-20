package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqUrlScan 恶意网址检测接口
// alibaba.security.jaq.url.scan
//
// url扫描接口
func AlibabaSecurityJaqUrlScan(clt *core.SDKClient, req *security.AlibabaSecurityJaqUrlScanAPIRequest, resp *security.AlibabaSecurityJaqUrlScanAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
