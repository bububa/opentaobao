package feedflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemCrowdPage 分页查询单品单元下人群列表
// taobao.feedflow.item.crowd.page
//
// 分页查询单品单元下人群列表
func TaobaoFeedflowItemCrowdPage(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCrowdPageAPIRequest, resp *feedflow.TaobaoFeedflowItemCrowdPageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
