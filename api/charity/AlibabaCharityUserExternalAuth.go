package charity

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCharityUserExternalAuth 外部用户授权
// alibaba.charity.user.external.auth
//
// 外部用户授权
func AlibabaCharityUserExternalAuth(ctx context.Context, clt *core.SDKClient, req *charity.AlibabaCharityUserExternalAuthAPIRequest, resp *charity.AlibabaCharityUserExternalAuthAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
