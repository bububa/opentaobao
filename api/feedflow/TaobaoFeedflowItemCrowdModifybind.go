package feedflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemCrowdModifybind 修改人群出价或状态
// taobao.feedflow.item.crowd.modifybind
//
// 修改人群出价或状态
func TaobaoFeedflowItemCrowdModifybind(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCrowdModifybindAPIRequest, resp *feedflow.TaobaoFeedflowItemCrowdModifybindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
