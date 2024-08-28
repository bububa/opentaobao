package security

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqAppRiskScanbatch 应用风险扫描批量提交接口
// alibaba.security.jaq.app.risk.scanbatch
//
// 批量提交应用进行风险扫描(含漏洞扫描、恶意代码检测),扫描完成后可通过对应的查询接口查询扫描结果
func AlibabaSecurityJaqAppRiskScanbatch(ctx context.Context, clt *core.SDKClient, req *security.AlibabaSecurityJaqAppRiskScanbatchAPIRequest, resp *security.AlibabaSecurityJaqAppRiskScanbatchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
