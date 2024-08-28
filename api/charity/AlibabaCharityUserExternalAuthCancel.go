package charity

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCharityUserExternalAuthCancel 取消外部授权
// alibaba.charity.user.external.auth.cancel
//
// 取消外部授权
func AlibabaCharityUserExternalAuthCancel(ctx context.Context, clt *core.SDKClient, req *charity.AlibabaCharityUserExternalAuthCancelAPIRequest, resp *charity.AlibabaCharityUserExternalAuthCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
