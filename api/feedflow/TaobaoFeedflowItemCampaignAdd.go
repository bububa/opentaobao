package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemcampaignadd 信息流增加推广计划
// taobao.feedflow.item.campaign.add
//
// 信息流增加推广计划
func Taobaofeedflowitemcampaignadd(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemcampaignaddAPIRequest, session string) (*feedflow.TaobaofeedflowitemcampaignaddAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemcampaignaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
