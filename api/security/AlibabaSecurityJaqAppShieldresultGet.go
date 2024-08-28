package security

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqAppShieldresultGet 用户查询加固结果
// alibaba.security.jaq.app.shieldresult.get
//
// 用户通过alibaba.security.jaq.app.shield接口提交应用加固后,通过该接口查询加固结果,下载加固包
func AlibabaSecurityJaqAppShieldresultGet(ctx context.Context, clt *core.SDKClient, req *security.AlibabaSecurityJaqAppShieldresultGetAPIRequest, resp *security.AlibabaSecurityJaqAppShieldresultGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
