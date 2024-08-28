package feedflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemCrowdRptdailylist 定向分日数据查询
// taobao.feedflow.item.crowd.rptdailylist
//
// 定向分日数据查询
func TaobaoFeedflowItemCrowdRptdailylist(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCrowdRptdailylistAPIRequest, resp *feedflow.TaobaoFeedflowItemCrowdRptdailylistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
