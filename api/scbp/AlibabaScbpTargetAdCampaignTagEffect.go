package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpTargetAdCampaignTagEffect 定向推广-获取推广计划定向效果数据
// alibaba.scbp.target.ad.campaign.tag.effect
//
// 定向推广-获取推广计划定向效果数据
func AlibabaScbpTargetAdCampaignTagEffect(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpTargetAdCampaignTagEffectAPIRequest, resp *scbp.AlibabaScbpTargetAdCampaignTagEffectAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
