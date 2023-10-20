package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemcampaignget 通过计划id查询计划
// taobao.feedflow.item.campaign.get
//
// 通过计划id查询计划
func Taobaofeedflowitemcampaignget(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemcampaigngetAPIRequest, session string) (*feedflow.TaobaofeedflowitemcampaigngetAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemcampaigngetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
