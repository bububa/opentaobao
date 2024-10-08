package security

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqAppRiskdetailbatchGet 应用风险详细信息批量查询接口
// alibaba.security.jaq.app.riskdetailbatch.get
//
// 用户通过alibaba.security.jaq.app.risk.scanbatch接口提交应用进行风险批量扫描后，用此接口批量获取风险详细信息,包含漏洞列表、恶意代码列表、仿冒应用列表等信息
func AlibabaSecurityJaqAppRiskdetailbatchGet(ctx context.Context, clt *core.SDKClient, req *security.AlibabaSecurityJaqAppRiskdetailbatchGetAPIRequest, resp *security.AlibabaSecurityJaqAppRiskdetailbatchGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
