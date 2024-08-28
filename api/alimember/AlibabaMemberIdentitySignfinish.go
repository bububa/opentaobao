package alimember

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimember"
)

// AlibabaMemberIdentitySignfinish 签约确认
// alibaba.member.identity.signfinish
//
// 签约确认
func AlibabaMemberIdentitySignfinish(ctx context.Context, clt *core.SDKClient, req *alimember.AlibabaMemberIdentitySignfinishAPIRequest, resp *alimember.AlibabaMemberIdentitySignfinishAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
