package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaSalestarCampaignBudgetUpdate 销量明星跟新预算相关接口
// taobao.simba.salestar.campaign.budget.update
//
// 更新一个推广计划的日限额
func TaobaoSimbaSalestarCampaignBudgetUpdate(clt *core.SDKClient, req *simba.TaobaoSimbaSalestarCampaignBudgetUpdateAPIRequest, resp *simba.TaobaoSimbaSalestarCampaignBudgetUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
