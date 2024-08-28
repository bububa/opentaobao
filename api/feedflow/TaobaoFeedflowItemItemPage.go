package feedflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemItemPage 信息流查看商品列表
// taobao.feedflow.item.item.page
//
// 信息流查看商品列表
func TaobaoFeedflowItemItemPage(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemItemPageAPIRequest, resp *feedflow.TaobaoFeedflowItemItemPageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
