package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaCampaignBudgetUpdate 更新一个推广计划的日限额
// taobao.simba.campaign.budget.update
//
// 更新一个推广计划的日限额
func TaobaoSimbaCampaignBudgetUpdate(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaCampaignBudgetUpdateAPIRequest, resp *simba.TaobaoSimbaCampaignBudgetUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
