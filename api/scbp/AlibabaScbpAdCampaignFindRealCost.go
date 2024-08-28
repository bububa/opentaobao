package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdCampaignFindRealCost 批量查询计划消耗数据
// alibaba.scbp.ad.campaign.find.real.cost
//
// 批量查询计划消耗数据
func AlibabaScbpAdCampaignFindRealCost(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdCampaignFindRealCostAPIRequest, resp *scbp.AlibabaScbpAdCampaignFindRealCostAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
