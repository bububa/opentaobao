package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdCampaignFindCampaignEffect 批量查询计划效果数据
// alibaba.scbp.ad.campaign.find.campaign.effect
//
// 批量查询计划效果数据
func AlibabaScbpAdCampaignFindCampaignEffect(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdCampaignFindCampaignEffectAPIRequest, resp *scbp.AlibabaScbpAdCampaignFindCampaignEffectAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
