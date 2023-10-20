package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemcampaignpage 批量查询计划列表
// taobao.feedflow.item.campaign.page
//
// 批量查询计划列表
func Taobaofeedflowitemcampaignpage(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemcampaignpageAPIRequest, session string) (*feedflow.TaobaofeedflowitemcampaignpageAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemcampaignpageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
