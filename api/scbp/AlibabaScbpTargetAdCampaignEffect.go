package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpTargetAdCampaignEffect 定向推广-获取计划维度推广效果
// alibaba.scbp.target.ad.campaign.effect
//
// 定向推广-获取计划维度推广效果
func AlibabaScbpTargetAdCampaignEffect(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpTargetAdCampaignEffectAPIRequest, resp *scbp.AlibabaScbpTargetAdCampaignEffectAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
