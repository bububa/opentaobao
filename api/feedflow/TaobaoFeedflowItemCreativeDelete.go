package feedflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemCreativeDelete 信息流删除创意
// taobao.feedflow.item.creative.delete
//
// 信息流删除创意
func TaobaoFeedflowItemCreativeDelete(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCreativeDeleteAPIRequest, resp *feedflow.TaobaoFeedflowItemCreativeDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
