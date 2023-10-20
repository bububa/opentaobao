package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemcampaignrptdailylist 推广计划分日数据查询
// taobao.feedflow.item.campaign.rptdailylist
//
// 推广计划分日数据查询
func Taobaofeedflowitemcampaignrptdailylist(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemcampaignrptdailylistAPIRequest, session string) (*feedflow.TaobaofeedflowitemcampaignrptdailylistAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemcampaignrptdailylistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
