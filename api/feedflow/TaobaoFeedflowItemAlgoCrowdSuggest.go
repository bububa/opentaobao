package feedflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemAlgoCrowdSuggest 单品人群建议出价
// taobao.feedflow.item.algo.crowd.suggest
//
// 给超级推荐的广告主查看建议出价
func TaobaoFeedflowItemAlgoCrowdSuggest(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAlgoCrowdSuggestAPIRequest, resp *feedflow.TaobaoFeedflowItemAlgoCrowdSuggestAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
