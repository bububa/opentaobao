package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemcampaigndelete 删除计划
// taobao.feedflow.item.campaign.delete
//
// 删除计划
func Taobaofeedflowitemcampaigndelete(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemcampaigndeleteAPIRequest, session string) (*feedflow.TaobaofeedflowitemcampaigndeleteAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemcampaigndeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
