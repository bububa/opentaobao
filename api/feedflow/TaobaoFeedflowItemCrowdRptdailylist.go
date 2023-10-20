package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemCrowdRptdailylist 定向分日数据查询
// taobao.feedflow.item.crowd.rptdailylist
//
// 定向分日数据查询
func TaobaoFeedflowItemCrowdRptdailylist(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCrowdRptdailylistAPIRequest, resp *feedflow.TaobaoFeedflowItemCrowdRptdailylistAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
