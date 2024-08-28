package security

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// AlibabaSecurityJaqAfsCheck 反欺诈二次验证接口
// alibaba.security.jaq.afs.check
//
// 反欺诈二次验证接口
func AlibabaSecurityJaqAfsCheck(ctx context.Context, clt *core.SDKClient, req *security.AlibabaSecurityJaqAfsCheckAPIRequest, resp *security.AlibabaSecurityJaqAfsCheckAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
