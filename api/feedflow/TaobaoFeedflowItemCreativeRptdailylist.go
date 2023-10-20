package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemCreativeRptdailylist 创意分日数据查询
// taobao.feedflow.item.creative.rptdailylist
//
// 创意分日数据查询
func TaobaoFeedflowItemCreativeRptdailylist(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCreativeRptdailylistAPIRequest, resp *feedflow.TaobaoFeedflowItemCreativeRptdailylistAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
