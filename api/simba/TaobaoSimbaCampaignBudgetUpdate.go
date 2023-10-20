package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaCampaignBudgetUpdate 更新一个推广计划的日限额
// taobao.simba.campaign.budget.update
//
// 更新一个推广计划的日限额
func TaobaoSimbaCampaignBudgetUpdate(clt *core.SDKClient, req *simba.TaobaoSimbaCampaignBudgetUpdateAPIRequest, resp *simba.TaobaoSimbaCampaignBudgetUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
