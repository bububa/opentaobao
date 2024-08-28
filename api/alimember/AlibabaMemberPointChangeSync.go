package alimember

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimember"
)

// AlibabaMemberPointChangeSync 成长值/积分变更记录同步
// alibaba.member.point.change.sync
//
// 成长值/积分变更记录同步
func AlibabaMemberPointChangeSync(ctx context.Context, clt *core.SDKClient, req *alimember.AlibabaMemberPointChangeSyncAPIRequest, resp *alimember.AlibabaMemberPointChangeSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
