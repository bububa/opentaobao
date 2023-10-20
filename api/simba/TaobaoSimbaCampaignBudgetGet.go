package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaCampaignBudgetGet 取得一个推广计划的日限额
// taobao.simba.campaign.budget.get
//
// 取得一个推广计划的日限额
func TaobaoSimbaCampaignBudgetGet(clt *core.SDKClient, req *simba.TaobaoSimbaCampaignBudgetGetAPIRequest, resp *simba.TaobaoSimbaCampaignBudgetGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
