package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemCampaignRpthourlist 超级推荐【商品推广】计划分时报表查询
// taobao.feedflow.item.campaign.rpthourlist
//
// 广告主推广计划分时数据查询，支持广告主查询最近90天内某一天的计划维度分时报表数据
func TaobaoFeedflowItemCampaignRpthourlist(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCampaignRpthourlistAPIRequest, resp *feedflow.TaobaoFeedflowItemCampaignRpthourlistAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
