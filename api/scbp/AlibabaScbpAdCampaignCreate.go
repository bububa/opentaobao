package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdCampaignCreate 创建计划
// alibaba.scbp.ad.campaign.create
//
// 创建计划
func AlibabaScbpAdCampaignCreate(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdCampaignCreateAPIRequest, resp *scbp.AlibabaScbpAdCampaignCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
