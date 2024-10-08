package alimember

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimember"
)

// AlibabaMemberIdentitySync 会员身份信息同步
// alibaba.member.identity.sync
//
// 会员身份信息同步
func AlibabaMemberIdentitySync(ctx context.Context, clt *core.SDKClient, req *alimember.AlibabaMemberIdentitySyncAPIRequest, resp *alimember.AlibabaMemberIdentitySyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
