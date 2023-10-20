package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemcampaignmodify 信息流修改计划
// taobao.feedflow.item.campaign.modify
//
// 信息流修改计划
func Taobaofeedflowitemcampaignmodify(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemcampaignmodifyAPIRequest, session string) (*feedflow.TaobaofeedflowitemcampaignmodifyAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemcampaignmodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
