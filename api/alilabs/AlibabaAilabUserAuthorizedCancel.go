package alilabs

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// AlibabaAilabUserAuthorizedCancel 取消账号授权
// alibaba.ailab.user.authorized.cancel
//
// 三方用户取消授权给天猫精灵用户
func AlibabaAilabUserAuthorizedCancel(ctx context.Context, clt *core.SDKClient, req *alilabs.AlibabaAilabUserAuthorizedCancelAPIRequest, resp *alilabs.AlibabaAilabUserAuthorizedCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
