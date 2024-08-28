package alimember

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimember"
)

// AlibabaMemberSync 会员信息同步
// alibaba.member.sync
//
// 会员信息同步
func AlibabaMemberSync(ctx context.Context, clt *core.SDKClient, req *alimember.AlibabaMemberSyncAPIRequest, resp *alimember.AlibabaMemberSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
