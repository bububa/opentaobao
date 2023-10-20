package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemcampaignrpthourlist 超级推荐【商品推广】计划分时报表查询
// taobao.feedflow.item.campaign.rpthourlist
//
// 广告主推广计划分时数据查询，支持广告主查询最近90天内某一天的计划维度分时报表数据
func Taobaofeedflowitemcampaignrpthourlist(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemcampaignrpthourlistAPIRequest, session string) (*feedflow.TaobaofeedflowitemcampaignrpthourlistAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemcampaignrpthourlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
