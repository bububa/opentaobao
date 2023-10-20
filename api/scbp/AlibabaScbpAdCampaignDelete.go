package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdCampaignDelete 删除计划
// alibaba.scbp.ad.campaign.delete
//
// 删除计划
func AlibabaScbpAdCampaignDelete(clt *core.SDKClient, req *scbp.AlibabaScbpAdCampaignDeleteAPIRequest, resp *scbp.AlibabaScbpAdCampaignDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
