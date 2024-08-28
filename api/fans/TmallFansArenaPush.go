package fans

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fans"
)

// TmallFansArenaPush 消息推送
// tmall.fans.arena.push
//
// 超级擂台消息推送
func TmallFansArenaPush(ctx context.Context, clt *core.SDKClient, req *fans.TmallFansArenaPushAPIRequest, resp *fans.TmallFansArenaPushAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
