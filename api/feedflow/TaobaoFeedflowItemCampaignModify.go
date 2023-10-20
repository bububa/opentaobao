package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemCampaignModify 信息流修改计划
// taobao.feedflow.item.campaign.modify
//
// 信息流修改计划
func TaobaoFeedflowItemCampaignModify(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCampaignModifyAPIRequest, resp *feedflow.TaobaoFeedflowItemCampaignModifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
