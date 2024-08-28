package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdCampaignDelete 删除计划
// alibaba.scbp.ad.campaign.delete
//
// 删除计划
func AlibabaScbpAdCampaignDelete(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdCampaignDeleteAPIRequest, resp *scbp.AlibabaScbpAdCampaignDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
