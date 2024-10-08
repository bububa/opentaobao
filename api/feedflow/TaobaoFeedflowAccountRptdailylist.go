package feedflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowAccountRptdailylist 获取广告主分日数据
// taobao.feedflow.account.rptdailylist
//
// 获取广告主分日数据
func TaobaoFeedflowAccountRptdailylist(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowAccountRptdailylistAPIRequest, resp *feedflow.TaobaoFeedflowAccountRptdailylistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
