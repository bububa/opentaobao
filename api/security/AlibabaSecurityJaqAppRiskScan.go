package security

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqAppRiskScan 应用风险扫描提交接口
// alibaba.security.jaq.app.risk.scan
//
// 提交应用进行风险扫描(含漏洞扫描、恶意代码检测、仿冒监测),扫描完成后可通过对应的查询接口查询扫描结果
func AlibabaSecurityJaqAppRiskScan(ctx context.Context, clt *core.SDKClient, req *security.AlibabaSecurityJaqAppRiskScanAPIRequest, resp *security.AlibabaSecurityJaqAppRiskScanAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
