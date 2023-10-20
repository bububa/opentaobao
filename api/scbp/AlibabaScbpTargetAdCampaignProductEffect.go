package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpTargetAdCampaignProductEffect 定向推广-获取计划中产品推广效果
// alibaba.scbp.target.ad.campaign.product.effect
//
// 定向推广-获取计划中产品推广效果
func AlibabaScbpTargetAdCampaignProductEffect(clt *core.SDKClient, req *scbp.AlibabaScbpTargetAdCampaignProductEffectAPIRequest, resp *scbp.AlibabaScbpTargetAdCampaignProductEffectAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
