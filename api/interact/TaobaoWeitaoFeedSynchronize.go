package interact

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// TaobaoWeitaoFeedSynchronize 推广淘小铺isv 活动到微淘feed
// taobao.weitao.feed.synchronize
//
// 推广淘小铺isv 活动到微淘feed
func TaobaoWeitaoFeedSynchronize(ctx context.Context, clt *core.SDKClient, req *interact.TaobaoWeitaoFeedSynchronizeAPIRequest, resp *interact.TaobaoWeitaoFeedSynchronizeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
