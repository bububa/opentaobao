package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdCampaignFindCampaignPage 分页查询计划
// alibaba.scbp.ad.campaign.find.campaign.page
//
// 分页查询计划
func AlibabaScbpAdCampaignFindCampaignPage(clt *core.SDKClient, req *scbp.AlibabaScbpAdCampaignFindCampaignPageAPIRequest, resp *scbp.AlibabaScbpAdCampaignFindCampaignPageAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
