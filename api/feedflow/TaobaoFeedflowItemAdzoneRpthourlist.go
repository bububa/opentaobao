package feedflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemAdzoneRpthourlist 超级推荐【商品推广】资源位分时报表查询
// taobao.feedflow.item.adzone.rpthourlist
//
// 广告主资源包分时数据查询，支持广告主查询最近90天内某一天的资源包维度分时报表数据
func TaobaoFeedflowItemAdzoneRpthourlist(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAdzoneRpthourlistAPIRequest, resp *feedflow.TaobaoFeedflowItemAdzoneRpthourlistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
