package feedflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemCrowdDelete 删除单品人群
// taobao.feedflow.item.crowd.delete
//
// 删除单品人群
func TaobaoFeedflowItemCrowdDelete(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCrowdDeleteAPIRequest, resp *feedflow.TaobaoFeedflowItemCrowdDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
