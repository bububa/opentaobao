package alimember

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimember"
)

// AlibabaMemberIdentityRescindfinish 取消确认
// alibaba.member.identity.rescindfinish
//
// 取消确认
func AlibabaMemberIdentityRescindfinish(ctx context.Context, clt *core.SDKClient, req *alimember.AlibabaMemberIdentityRescindfinishAPIRequest, resp *alimember.AlibabaMemberIdentityRescindfinishAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
