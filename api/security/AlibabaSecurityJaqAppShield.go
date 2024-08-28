package security

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqAppShield 应用加固接口
// alibaba.security.jaq.app.shield
//
// 提交应用进行应用加固,加固后需通过alibaba.security.jaq.app.shieldresult.get接口查询加固结果
func AlibabaSecurityJaqAppShield(ctx context.Context, clt *core.SDKClient, req *security.AlibabaSecurityJaqAppShieldAPIRequest, resp *security.AlibabaSecurityJaqAppShieldAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
