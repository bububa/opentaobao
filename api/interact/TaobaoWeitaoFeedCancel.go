package interact

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// TaobaoWeitaoFeedCancel 取消广播在timeline、广场中展示
// taobao.weitao.feed.cancel
//
// 取消广播在timeline和广场中的展示。
func TaobaoWeitaoFeedCancel(ctx context.Context, clt *core.SDKClient, req *interact.TaobaoWeitaoFeedCancelAPIRequest, resp *interact.TaobaoWeitaoFeedCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
