package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdCampaignUpdate 修改计划
// alibaba.scbp.ad.campaign.update
//
// 修改计划
func AlibabaScbpAdCampaignUpdate(clt *core.SDKClient, req *scbp.AlibabaScbpAdCampaignUpdateAPIRequest, resp *scbp.AlibabaScbpAdCampaignUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
