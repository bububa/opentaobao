package feedflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemCreativeRpthourlist 超级推荐【商品推广】创意分时报表查询
// taobao.feedflow.item.creative.rpthourlist
//
// 创意分时数据查询，支持广告主查询最近90天内某一天的创意维度分时报表数据
func TaobaoFeedflowItemCreativeRpthourlist(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCreativeRpthourlistAPIRequest, resp *feedflow.TaobaoFeedflowItemCreativeRpthourlistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
