package security

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqUrlScan 恶意网址检测接口
// alibaba.security.jaq.url.scan
//
// url扫描接口
func AlibabaSecurityJaqUrlScan(ctx context.Context, clt *core.SDKClient, req *security.AlibabaSecurityJaqUrlScanAPIRequest, resp *security.AlibabaSecurityJaqUrlScanAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
